package provider

import (
	"testing"
)

func TestCreation(t *testing.T) {
	provider := New()
	if provider == nil {
		t.Fail()
	}
}

func TestGetURL(t *testing.T) {
	provider := New()

	testURL := provider.getURL("/signin/")
	if testURL != "https://njal.la/signin/" {
		t.Fail()
	}
}

func TestDomainURL(t *testing.T) {
	provider := New()

	testURL := provider.getDomainURL("mydomain.com")
	if testURL != "https://njal.la/domains/mydomain.com/" {
		t.Fail()
	}
}
