package provider

import "testing"

func TestCreation(t *testing.T) {
	provider := New("test.com")
	if provider == nil {
		t.Fail()
	}
}

func TestGetURL(t *testing.T) {
	provider := New("test.com")
	testURL := provider.getURL("/test")
	if testURL != "test.com/test" {
		t.Fail()
	}
}

func TestDomainURL(t *testing.T) {
	provider := New("test.com")
	if provider.getDomainURL() != "test.com" {
		t.Fail()
	}
}
