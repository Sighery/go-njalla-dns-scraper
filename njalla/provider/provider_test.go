package provider

import (
	"fmt"
	"testing"
)

func TestCreation(t *testing.T) {
	provider, err := New()
	if err != nil {
		t.Errorf("%s", err)
	} else if provider == nil || provider.BaseURL != "https://njal.la" {
		t.Fail()
	}
}

func TestGetURL(t *testing.T) {
	provider, err := New()
	if err != nil {
		t.Errorf("%s", err)
	}

	testURL := provider.getURL("/signin/")
	if testURL != "https://njal.la/signin/" {
		t.Fail()
	}
}

func TestDomainURL(t *testing.T) {
	provider, err := New()
	if err != nil {
		t.Errorf("%s", err)
	}

	testURL := provider.getDomainURL("mydomain.com")
	if testURL != "https://njal.la/domains/mydomain.com/" {
		t.Fail()
	}
}

func TestGetDomains(t *testing.T) {
	provider, _ := New()
	provider.Login("email", `password`)
	domains, _ := provider.GetDomains()
	fmt.Println(domains)
}
