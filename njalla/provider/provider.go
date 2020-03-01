package provider

import (
	"fmt"
)

// Provider struct
type Provider struct {
	BaseURL string
}

// New constructor creates a given domain provider with its base URL
func New() *Provider {

	return &Provider{
		BaseURL: "https://njal.la",
	}
}

// getURL returns a given domain provider URL, such as https://njal.la/signin/
func (p *Provider) GetURL(path string) string {
	return fmt.Sprintf("%s%s", p.BaseURL, path)
}

// getDomainURL returns the provider URL for a given domain configuration,
// such as https://njal.la/domains/mydomain.com/
func (p *Provider) GetDomainURL(domain string) string {
	return fmt.Sprintf("%s%s/", p.GetURL("/domains/"), domain)
}
