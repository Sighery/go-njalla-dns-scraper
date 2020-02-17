package provider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Sighery/go-njalla-dns-scraper/njalla/records"
	"github.com/google/go-querystring/query"
)

// Provider struct
type Provider struct {
	BaseURL string
	jar     *cookiejar.Jar
	client  http.Client
	Records records.Records
}

// New constructor creates a given domain provider with its base URL
func New() (*Provider, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	return &Provider{
		BaseURL: "https://njal.la",
		jar:     jar,
		client:  http.Client{Jar: jar},
		Records: records.Records{},
	}, nil
}

// getURL returns a given domain provider URL, such as https://njal.la/signin/
func (p *Provider) getURL(path string) string {
	return fmt.Sprintf("%s%s", p.BaseURL, path)
}

// getDomainURL returns the provider URL for a given domain configuration,
// such as https://njal.la/domains/mydomain.com/
func (p *Provider) getDomainURL(domain string) string {
	return fmt.Sprintf("%s%s/", p.getURL("/domains/"), domain)
}

// initialize does a first time set up needed to fetch cookies and CSRF token
func (p *Provider) initialize() (string, error) {
	resp, err := p.client.Get(p.getURL("/signin/"))
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	value, exists := doc.Find("input[name=\"csrfmiddlewaretoken\"]").First().
		Attr("value")
	if !exists {
		return "", fmt.Errorf("Couldn't find input with CSRF token")
	}

	return value, nil
}

// Login logs a given user in
func (p *Provider) Login(email, password string) error {
	csrf, err := p.initialize()
	if err != nil {
		return err
	}

	values := url.Values{}
	values.Set("csrfmiddlewaretoken", csrf)
	values.Set("email", email)
	values.Set("password", password)

	resp, respErr := postForm(p.client, p.getURL("/signin/"), values)
	if respErr != nil {
		return respErr
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Login failed with status code %d", resp.StatusCode)
	}

	return nil
}

// GetDomains returns an array of available domains in your Njalla account
func (p *Provider) GetDomains() ([]string, error) {
	_, err := getCSRFToken(p.jar, p.BaseURL)
	if err != nil {
		return nil, err
	}

	resp, respErr := p.client.Get(p.getURL("/domains/"))
	if respErr != nil {
		return nil, respErr
	}

	doc, docErr := goquery.NewDocumentFromReader(resp.Body)
	if docErr != nil {
		return nil, docErr
	}

	query := "/domains/"
	startIndex := len(query)
	domains := make([]string, 0)

	doc.Find(".table a[href^=\"/domains/\"]:contains(Manage)").
		Each(func(i int, s *goquery.Selection) {
			href, _ := s.Attr("href")
			domains = append(domains, href[startIndex:len(href)-1])
		})

	return domains, nil
}

// GetRecords returns Records with all the available records for a domain
func (p *Provider) GetRecords(domain string) (records.Records, error) {
	_, err := getCSRFToken(p.jar, p.BaseURL)
	if err != nil {
		return nil, err
	}

	resp, respErr := p.client.Get(p.getDomainURL(domain))
	if respErr != nil {
		return nil, respErr
	}

	doc, docErr := goquery.NewDocumentFromReader(resp.Body)
	if docErr != nil {
		return nil, docErr
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

	var r records.Records
	jsonErr := json.Unmarshal([]byte(match), &r)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return r, nil
}

// AddRecord creates a new record in Njalla. It accepts only those record
// types defined in records/records
func (p *Provider) AddRecord(domain string, record records.Record) error {
	csrftoken, loginErr := getCSRFToken(p.jar, p.BaseURL)
	if loginErr != nil {
		return loginErr
	}

	values, vErr := query.Values(record)
	if vErr != nil {
		return vErr
	}

	values.Set("action", "add")
	values.Set("csrfmiddlewaretoken", csrftoken)

	resp, respErr := postForm(p.client, p.getDomainURL(domain), values)
	if respErr != nil {
		return respErr
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf(
			"Adding record failed with status code %d", resp.StatusCode,
		)
	}

	return nil
}

func postForm(client http.Client, path string, data url.Values) (*http.Response, error) {
	req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Referer", path)
	return client.Do(req)
}

func getCSRFToken(jar *cookiejar.Jar, URL string) (string, error) {
	// This function can be used to check if the program is logged in. If we
	// aren't, then the cookiejar won't have this csrf token set.
	parsedURL, parsedErr := url.Parse(URL)
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
		return "", fmt.Errorf("Not logged in")
	}

	return csrftoken, nil
}