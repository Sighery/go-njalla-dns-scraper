package provider

import "fmt"

// Provider struct
type Provider struct {
	URL string
}

// Providers interface implemented for providers
type Providers interface {
	getUrl()
	getDomainURL()
	login()
}

// New constructor creates a given domain provider with its base URL
func New(url string) *Provider {
	return &Provider{
		URL: url,
	}
}

// getURL returns a given domain provider URL, such as https://njal.la/signin/
func (p *Provider) getURL(path string) string {
	return fmt.Sprintf("%s%s", p.URL, path)
}

// getDomainURL returns the provider URL for a given domain configuration,
// such as https://njal.la/domains/mydomain.com/
func (p *Provider) getDomainURL(domain string) string {
	return fmt.Sprintf("%s%s/", p.getURL("/domains/"), domain)
}
