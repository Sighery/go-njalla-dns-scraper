package records

import (
	"encoding/json"
	"fmt"
)

type Records []record

func (r *Records) UnmarshalJSON(data []byte) error {
	// This just splits up the JSON array into the raw JSON for each object
	var raw []json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	for _, x := range raw {
		// Unmarshal into a map to check the "type" field
		var obj map[string]interface{}
		err := json.Unmarshal(x, &obj)
		if err != nil {
			return err
		}

		recordType := ""
		if t, ok := obj["type"].(string); ok {
			recordType = t
		} else {
			return fmt.Errorf("Record doesn't have field type: %v", obj)
		}

		// Unmarshal again into the correct type
		var actual record
		switch recordType {
		case "A":
			actual = &RecordA{}
		case "AAAA":
			actual = &RecordAAAA{}
		case "CNAME":
			actual = &RecordCNAME{}
		case "MX":
			actual = &RecordMX{}
		case "TXT":
			actual = &RecordTXT{}
		case "SRV":
			actual = &RecordSRV{}
		case "CAA":
			actual = &RecordCAA{}
		case "PTR":
			actual = &RecordPTR{}
		case "NS":
			actual = &RecordNS{}
		case "TLSA":
			actual = &RecordTLSA{}
		case "Redirect":
			actual = &RecordRedirect{}
		case "Dynamic":
			actual = &RecordDynamic{}
		case "SSHFP":
			actual = &RecordSSHFP{}
		default:
			return fmt.Errorf("Unknown record type: %s", recordType)
		}

		err = json.Unmarshal(x, actual)
		if err != nil {
			return err
		}
		*r = append(*r, actual)
	}

	return nil
}

func (r Records) String() string {
	representation := ""
	for _, record := range r {
		representation += fmt.Sprintf("%+v", record)
	}
	return representation
}

type record interface{}

// A record type
type RecordA struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

// AAAA record type
type RecordAAAA struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

// CNAME record type
type RecordCNAME struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

// RecordMX type
type RecordMX struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	TTL      int    `json:"ttl"`
	Priority int    `json:"prio"`
}

// TXT record type
type RecordTXT struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

// RecordSRV type
type RecordSRV struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	TTL      int    `json:"ttl"`
	Priority int    `json:"prio"`
	Weight   int    `json:"weight"`
	Port     int    `json:"port"`
}

// CAA record type
type RecordCAA struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

// PTR record type
type RecordPTR struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

// NS record type
type RecordNS struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

// TLSA record type
type RecordTLSA struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

// Redirect record type
type RecordRedirect struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	URL      string `json:"content"`
	Priority int    `json:"prio"`
}

// Dynamic record type
type RecordDynamic struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	TTL  int    `json:"ttl"`
}

// SSHFP record type
type RecordSSHFP struct {
	ID           int    `json:"id"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	TTL          int    `json:"ttl"`
	SSHAlgorithm int    `json:"ssh_algorithm"`
	SSHType      int    `json:"ssh_type"`
	Content      string `json:"content"`
}
