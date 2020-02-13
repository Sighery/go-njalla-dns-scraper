package main

import (
	"errors"
	"fmt"

	"strings"

	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

var jar, _ = cookiejar.New(nil)
var client = http.Client{Jar: jar}

const baseURL string = "https://njal.la"

func getNjallaURL(path string) string {
	return fmt.Sprintf("%s%s", baseURL, path)
}
func getNjallaDomainURL(domain string) string {
	return fmt.Sprintf("%s%s/", getNjallaURL("/domains/"), domain)
}

func postForm(client http.Client, path string, data url.Values) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Referer", path)
	return client.Do(req)
}

func getCSRFToken(jar *cookiejar.Jar) (string, error) {
	// This function can be used to check if the program is logged in. If we
	// aren't, then the cookiejar won't have this csrf token set.

	parsedURL, parsedErr := url.Parse(baseURL)
	if parsedErr != nil {
		return "", parsedErr
	}

	csrftoken := ""

	for _, cookie := range jar.Cookies(parsedURL) {
		if cookie.Name == "csrftoken" {
			csrftoken = cookie.Value
			break
		}
	}

	if len(csrftoken) < 0 {
		return "", errors.New("not logged in")
	}

	return csrftoken, nil
}

type Record struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	TTL      int    `json:"ttl"`
	Priority *int   `json:"prio,omitempty"`
}

func (r Record) String() string {
	if r.Priority != nil {
		return fmt.Sprintf(
			"Record{ID:%d, Type:%s, Name:%s, Content:%s, TTL:%d, Priority:%v}",
			r.ID, r.Type, r.Name, r.Content, r.TTL, *r.Priority,
		)
	}

	return fmt.Sprintf(
		"Record{ID:%d, Type:%s, Name:%s, Content:%s, TTL:%d, Priority:%v}",
		r.ID, r.Type, r.Name, r.Content, r.TTL, r.Priority,
	)
}

func addGenericRecord(domain string, record url.Values) error {
	csrftoken, loginErr := getCSRFToken(jar)
	if loginErr != nil {
		return loginErr
	}

	record.Set("action", "add")
	record.Set("csrfmiddlewaretoken", csrftoken)

	_, err := postForm(client, getNjallaDomainURL(domain), record)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Example usage. This won't work without proper credentials
	loginErr := login("email", "password")
	if loginErr != nil {
		panic(loginErr)
	}

	domains, domainsErr := getDomains()
	if domainsErr != nil {
		panic(domainsErr)
	}

	fmt.Println("Available domains:", domains)

	var records []Record
	records, err := getRecords("domain.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Available records for domain %s:", "domain.com"))
	for _, record := range records {
		fmt.Println(record)
	}

	// Adding some records. You can check it worked by fetching available
	// records again afterwards
	addTXTRecord("domain.com", "TXT", "@", "TEST", 10800)
	addARecord("domain.com", "A", "www", "1.2.3.4", 10800)
}

func initialize() (string, error) {
	resp, err := client.Get(getNjallaURL("/signin/"))

	if err != nil {
		return "", err
	}

	z := html.NewTokenizer(resp.Body)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done. Function should have ended
			// before reaching here
			return "", errors.New("Couldn't find CSRF token")
		case tt == html.SelfClosingTagToken:
			t := z.Token()

			isInput := t.Data == "input"
			if isInput {
				for _, a := range t.Attr {
					if a.Key == "name" && a.Val == "csrfmiddlewaretoken" {
						for _, a := range t.Attr {
							if a.Key == "value" {
								return a.Val, nil
							}
						}
					}
				}
			}
		}
	}
}

func login(email, password string) error {
	csrf, err := initialize()
	if err != nil {
		return err
	}

	values := url.Values{}
	values.Set("csrfmiddlewaretoken", csrf)
	values.Set("email", email)
	values.Set("password", password)

	_, respErr := postForm(client, getNjallaURL("/signin/"), values)
	if respErr != nil {
		return respErr
	}

	return nil
}

func getDomains() ([]string, error) {
	_, err := getCSRFToken(jar)
	if err != nil {
		return nil, err
	}

	resp, err := client.Get(getNjallaURL("/domains"))
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	query := "/domains/"
	startIndex := len(query)
	domains := make([]string, 1)

	doc.Find(".table a[href^=\"/domains/\"]:contains(Manage)").
		Each(func(i int, s *goquery.Selection) {
			href, _ := s.Attr("href")
			domains = append(domains, href[startIndex:len(href)-1])
		})

	return domains, nil
}

func getRecords(domain string) ([]Record, error) {
	_, err := getCSRFToken(jar)
	if err != nil {
		return nil, err
	}

	resp, err := client.Get(getNjallaDomainURL(domain))
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	query := "var records = "
	finish := "];\n"
	var match string

	doc.Find(fmt.Sprintf("script:contains(\"%s\")", query)).
		Each(func(i int, s *goquery.Selection) {
			text := s.Text()

			startIndex := strings.Index(text, query)
			if startIndex == -1 {
				return
			}

			endIndex := strings.Index(text[startIndex:], finish)
			if endIndex == -1 {
				return
			}

			// Shift end index to include the `]`
			endIndex += startIndex + 1
			// Shift start index to fetch starting at `[`
			startIndex += len(query)

			match = text[startIndex:endIndex]
		})

	var records []Record
	json.Unmarshal([]byte(match), &records)
	return records, nil
}

func addTXTRecord(domain string, rtype string, name string, content string, ttl int) error {
	values := url.Values{}
	values.Set("type", rtype)
	values.Set("name", name)
	values.Set("content", content)
	values.Set("ttl", fmt.Sprintf("%d", ttl))

	err := addGenericRecord(domain, values)
	if err != nil {
		return err
	}

	return nil
}

func addARecord(domain string, rtype string, name string, content string, ttl int) error {
	values := url.Values{}
	values.Set("type", rtype)
	values.Set("name", name)
	values.Set("content", content)
	values.Set("ttl", fmt.Sprintf("%d", ttl))

	err := addGenericRecord(domain, values)
	if err != nil {
		return err
	}

	return nil
}

func addAAAARecord(domain string, rtype string, name string, content string, ttl int) error {
	values := url.Values{}
	values.Set("type", rtype)
	values.Set("name", name)
	values.Set("content", content)
	values.Set("ttl", fmt.Sprintf("%d", ttl))

	err := addGenericRecord(domain, values)
	if err != nil {
		return err
	}

	return nil
}

func addMXRecord(domain string, rtype string, name string, content string, ttl int, priority int) error {
	values := url.Values{}
	values.Set("type", rtype)
	values.Set("name", name)
	values.Set("content", content)
	values.Set("ttl", fmt.Sprintf("%d", ttl))
	values.Set("priority", fmt.Sprintf("%d", priority))

	err := addGenericRecord(domain, values)
	if err != nil {
		return err
	}

	return nil
}
