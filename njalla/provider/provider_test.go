package provider

import (
	// "fmt"
	"testing"

	// "github.com/Sighery/go-njalla-dns-scraper/njalla/records"
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

// func TestUpdateDomain(t *testing.T) {
// 	provider, _ := New()
// 	provider.Login("email", `password`)
// 	records, _ := provider.GetRecords("sighery.com")
// 	for _, r := range records {
// 		if r.GetID() == 118747 {
// 			toChange := r.GetURLValues()
// 			toChange.Set("content", "TEST3")
// 			err := provider.UpdateRecord(
// 				"sighery.com", r.GetID(), toChange,
// 			)
// 			if err != nil {
// 				t.Error(err)
// 				return
// 			}
// 		}
// 	}
// }

// func TestRemoveDomain(t *testing.T) {
// 	provider, _ := New()
// 	provider.Login("email", `password`)
// 	err := provider.RemoveRecord("sighery.com", 118760)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestGetDomains(t *testing.T) {
// 	provider, _ := New()
// 	provider.Login("email", `password`)
// 	domains, _ := provider.GetDomains()
// 	fmt.Println(domains)
// }

// func TestGetRecords(t *testing.T) {
// 	provider, _ := New()
// 	provider.Login("email", `password`)
// 	records, _ := provider.GetRecords("sighery.com")
// 	for _, record := range records {
// 		fmt.Println(record)
// 	}
// 	fmt.Println(records)
// }

// func TestAddRecord(t *testing.T) {
// 	provider, _ := New()
// 	provider.Login("email", `password`)
// 	record := records.RecordTXT{
// 		Type: "TXT",
// 		Name: "test",
// 		Content: "TEST",
// 		TTL: 10800,
// 	}
// 	provider.AddRecord("sighery.com", record)
// 	fmt.Println("Yes")
// }
