package provider

import "testing"

func TestCreation(t *testing.T) {
	provider := New("https://njal.la")
	if provider == nil {
		t.Fail()
	}
}

func TestGetURL(t *testing.T) {
	provider := New("https://njal.la")
	testURL := provider.getURL("/signin/")
	if testURL != "https://njal.la/signin/" {
		t.Fail()
	}
}

func TestDomainURL(t *testing.T) {
	provider := New("https://njal.la")
	testURL := provider.getDomainURL("mydomain.com")
	if testURL != "https://njal.la/domains/mydomain.com/" {
		t.Fail()
	}
}
