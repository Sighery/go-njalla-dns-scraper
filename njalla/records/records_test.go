package records

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestValidJSONWorks(t *testing.T) {
	js := []byte(`[
	{
		"type": "A",
		"name": "@",
		"id": 1,
		"content": "1.1.1.1",
		"ttl": 10800
	},
	{
		"type": "AAAA",
		"name": "@",
		"id": 2,
		"content": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
		"ttl": 10800
	},
	{
		"type": "CNAME",
		"name": "@",
		"id": 3,
		"content": "Content",
		"ttl": 10800
	},
	{
		"type": "MX",
		"name": "@",
		"id": 4,
		"content": "mail.protonmail.ch",
		"ttl": 10800,
		"prio": 10
	},
	{
		"type": "TXT",
		"name": "@",
		"id": 5,
		"content": "v=spf1 a mx ?all",
		"ttl": 10800
	},
	{
		"type": "SRV",
		"name": "@",
		"id": 6,
		"content": "Content",
		"ttl": 10800,
		"prio": 10,
		"weight": 0,
		"port": 1234
	},
	{
		"type": "CAA",
		"name": "@",
		"id": 7,
		"content": "0 issue \"letsencrypt.org\"",
		"ttl": 10800
	},
	{
		"type": "PTR",
		"name": "@",
		"id": 8,
		"content": "Content",
		"ttl": 10800
	},
	{
		"type": "NS",
		"name": "Custom",
		"id": 9,
		"content": "3-get.njalla.fo",
		"ttl": 10800
	},
	{
		"type": "TLSA",
		"name": "@",
		"id": 10,
		"content": "3 1 1 BASE64==",
		"ttl": 10800
	},
	{
		"type": "Redirect",
		"name": "@",
		"id": 11,
		"content": "https://example.com",
		"prio": 301
	},
	{
		"type": "Dynamic",
		"name": "@",
		"id": 12,
		"ttl": 60
	},
	{
		"type": "SSHFP",
		"name": "@",
		"id": 13,
		"content": "Content",
		"ttl": 10800,
		"ssh_algorithm": 4,
		"ssh_type": 2
	}
]`)

	var r Records
	err := json.Unmarshal(js, &r)
	if err != nil {
		t.Fatalf("%q", err)
	}

	expected := Records{
		&RecordA{ID: 1, Type: "A", Name: "@", Content: "1.1.1.1", TTL: 10800},
		&RecordAAAA{
			ID: 2, Type: "AAAA", Name: "@",
			Content: "2001:0db8:85a3:0000:0000:8a2e:0370:7334", TTL: 10800,
		},
		&RecordCNAME{
			ID: 3, Type: "CNAME", Name: "@", Content: "Content", TTL: 10800,
		},
		&RecordMX{
			ID: 4, Type: "MX", Name: "@", Content: "mail.protonmail.ch",
			TTL: 10800, Priority: 10,
		},
		&RecordTXT{
			ID: 5, Type: "TXT", Name: "@", Content: "v=spf1 a mx ?all",
			TTL: 10800,
		},
		&RecordSRV{
			ID: 6, Type: "SRV", Name: "@", Content: "Content", TTL: 10800,
			Priority: 10, Weight: 0, Port: 1234,
		},
		&RecordCAA{
			ID: 7, Type: "CAA", Name: "@", TTL: 10800,
			Content: `0 issue "letsencrypt.org"`,
		},
		&RecordPTR{
			ID: 8, Type: "PTR", Name: "@", Content: "Content", TTL: 10800,
		},
		&RecordNS{
			ID: 9, Type: "NS", Name: "Custom", Content: "3-get.njalla.fo",
			TTL: 10800,
		},
		&RecordTLSA{
			ID: 10, Type: "TLSA", Name: "@", Content: "3 1 1 BASE64==",
			TTL: 10800,
		},
		&RecordRedirect{
			ID: 11, Type: "Redirect", Name: "@",
			URL: "https://example.com", RedirectType: 301,
		},
		&RecordDynamic{ID: 12, Type: "Dynamic", Name: "@", TTL: 60},
		&RecordSSHFP{
			ID: 13, Type: "SSHFP", Name: "@", Content: "Content", TTL: 10800,
			SSHAlgorithm: 4, SSHType: 2,
		},
	}

	if rLen, exLen := len(r), len(expected); rLen != exLen {
		t.Errorf(
			`Length of parsed [%d] and length of expected [%d] don't match
Parsed:
%+v

Expected:
%+v`,
			rLen, exLen, r, expected,
		)
	}

	for i, e := range expected {
		testname := fmt.Sprintf("Index: %d", i)

		t.Run(testname, func(t *testing.T) {
			if rRec := r[i]; !cmp.Equal(e, rRec) {
				t.Fatalf(
					"Record from parsed:\n%+v\ndoesn't match with expected:\n%+v",
					e, rRec,
				)
			}
		})
	}
}

func TestInvalidTypeFails(t *testing.T) {
	js := []byte(`[
	{
		"type": "MADEUP",
		"id": 0,
		"name": "fake",
		"content": "no"
	}
]`)

	var r Records
	err := json.Unmarshal(js, &r)

	if err == nil || !strings.Contains(err.Error(), "Unknown record type") {
		t.Fatalf("Test didn't fail as expected")
	}
}

func TestMissingTypeFails(t *testing.T) {
	js := []byte(`[
	{
		"id": 0,
		"name": "fake",
		"content": "no"
	}
]`)

	var r Records
	err := json.Unmarshal(js, &r)

	if err == nil || !strings.Contains(err.Error(), "doesn't have field type") {
		t.Fatalf("Test didn't fail as expected")
	}
}
