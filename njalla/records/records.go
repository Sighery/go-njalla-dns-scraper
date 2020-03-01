package records

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/Sighery/go-njalla-dns-scraper/njalla/structures"
)

// Records is an array of different record types that implement the Record
// interface
type Records []Record

// UnmarshalJSON customises the default unmarshal behaviour to parse into
// specific record types
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
		var actual Record
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
		s, _ := json.Marshal(record)
		representation += fmt.Sprintf("%s\n", s)
	}
	return representation
}

// Record is a common interface for all the specific record types
type Record interface {
	GetURLValues() url.Values
	GetID() int
}

// RecordA represents Njalla's A record
type RecordA struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

// GetURLValues converts struct fields back into provider suitable values
func (r RecordA) GetURLValues() url.Values {
	values := url.Values{}
	values.Set("id", fmt.Sprintf("%d", r.ID))
	values.Set("type", r.Type)
	values.Set("name", r.Name)
	values.Set("content", r.Content)
	values.Set("ttl", fmt.Sprintf("%d", r.TTL))
	return values
}

// GetID exposes the internal ID
func (r RecordA) GetID() int {
	return r.ID
}

// NewRecordA validates and creates a new RecordA
func NewRecordA(name string, content string, ttl int) (RecordA, error) {
	if ttlErr := checkValidTTL(ttl); ttlErr != nil {
		return RecordA{}, ttlErr
	}

	return RecordA{
		Type:    "A",
		Name:    name,
		Content: content,
		TTL:     ttl,
	}, nil
}

// RecordAAAA represents Njalla's AAAA record
type RecordAAAA struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

// GetURLValues converts struct fields back into provider suitable values
func (r RecordAAAA) GetURLValues() url.Values {
	values := url.Values{}
	values.Set("id", fmt.Sprintf("%d", r.ID))
	values.Set("type", r.Type)
	values.Set("name", r.Name)
	values.Set("content", r.Content)
	values.Set("ttl", fmt.Sprintf("%d", r.TTL))
	return values
}

// GetID exposes the internal ID
func (r RecordAAAA) GetID() int {
	return r.ID
}

// NewRecordAAAA validates and creates a new RecordAAAA
func NewRecordAAAA(name string, content string, ttl int) (RecordAAAA, error) {
	if ttlErr := checkValidTTL(ttl); ttlErr != nil {
		return RecordAAAA{}, ttlErr
	}

	return RecordAAAA{
		Type:    "AAAA",
		Name:    name,
		Content: content,
		TTL:     ttl,
	}, nil
}

// RecordCNAME represents Njalla's CNAME record
type RecordCNAME struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

// GetURLValues converts struct fields back into provider suitable values
func (r RecordCNAME) GetURLValues() url.Values {
	values := url.Values{}
	values.Set("id", fmt.Sprintf("%d", r.ID))
	values.Set("type", r.Type)
	values.Set("name", r.Name)
	values.Set("content", r.Content)
	values.Set("ttl", fmt.Sprintf("%d", r.TTL))
	return values
}

// GetID exposes the internal ID
func (r RecordCNAME) GetID() int {
	return r.ID
}

// NewRecordCNAME validates and creates a new RecordCNAME
func NewRecordCNAME(name string, content string, ttl int) (RecordCNAME, error) {
	if ttlErr := checkValidTTL(ttl); ttlErr != nil {
		return RecordCNAME{}, ttlErr
	}

	return RecordCNAME{
		Type:    "CNAME",
		Name:    name,
		Content: content,
		TTL:     ttl,
	}, nil
}

// RecordMX represents Njalla's MX record
type RecordMX struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	TTL      int    `json:"ttl"`
	Priority int    `json:"prio"`
}

// GetURLValues converts struct fields back into provider suitable values
func (r RecordMX) GetURLValues() url.Values {
	values := url.Values{}
	values.Set("id", fmt.Sprintf("%d", r.ID))
	values.Set("type", r.Type)
	values.Set("name", r.Name)
	values.Set("content", r.Content)
	values.Set("ttl", fmt.Sprintf("%d", r.TTL))
	values.Set("prio", fmt.Sprintf("%d", r.Priority))
	return values
}

// GetID exposes the internal ID
func (r RecordMX) GetID() int {
	return r.ID
}

// NewRecordMX validates and creates a new RecordMX
func NewRecordMX(name string, content string, ttl int, priority int) (RecordMX, error) {
	if ttlErr := checkValidTTL(ttl); ttlErr != nil {
		return RecordMX{}, ttlErr
	}

	if priErr := checkValidPriority(priority); priErr != nil {
		return RecordMX{}, priErr
	}

	return RecordMX{
		Type:     "MX",
		Name:     name,
		Content:  content,
		TTL:      ttl,
		Priority: priority,
	}, nil
}

// RecordTXT represents Njalla's TXT record
type RecordTXT struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

// GetURLValues converts struct fields back into provider suitable values
func (r RecordTXT) GetURLValues() url.Values {
	values := url.Values{}
	values.Set("id", fmt.Sprintf("%d", r.ID))
	values.Set("type", r.Type)
	values.Set("name", r.Name)
	values.Set("content", r.Content)
	values.Set("ttl", fmt.Sprintf("%d", r.TTL))
	return values
}

// GetID exposes the internal ID
func (r RecordTXT) GetID() int {
	return r.ID
}

// NewRecordTXT validates and creates a new RecordTXT
func NewRecordTXT(name string, content string, ttl int) (RecordTXT, error) {
	if ttlErr := checkValidTTL(ttl); ttlErr != nil {
		return RecordTXT{}, ttlErr
	}

	return RecordTXT{
		Type:    "TXT",
		Name:    name,
		Content: content,
		TTL:     ttl,
	}, nil
}

// RecordSRV represents Njalla's SRV record
type RecordSRV struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	TTL      int    `json:"ttl"`
	Priority int    `json:"prio"`
	Weight   uint   `json:"weight"`
	Port     uint   `json:"port"`
}

// GetURLValues converts struct fields back into provider suitable values
func (r RecordSRV) GetURLValues() url.Values {
	values := url.Values{}
	values.Set("id", fmt.Sprintf("%d", r.ID))
	values.Set("type", r.Type)
	values.Set("name", r.Name)
	values.Set("content", r.Content)
	values.Set("ttl", fmt.Sprintf("%d", r.TTL))
	values.Set("prio", fmt.Sprintf("%d", r.Priority))
	values.Set("weight", fmt.Sprintf("%d", r.Weight))
	values.Set("port", fmt.Sprintf("%d", r.Port))
	return values
}

// GetID exposes the internal ID
func (r RecordSRV) GetID() int {
	return r.ID
}

// NewRecordSRV validates and creates a new RecordSRV
func NewRecordSRV(
	name string, content string, ttl int, priority int, weight uint,
	port uint,
) (RecordSRV, error) {
	if ttlErr := checkValidTTL(ttl); ttlErr != nil {
		return RecordSRV{}, ttlErr
	}

	if priErr := checkValidPriority(priority); priErr != nil {
		return RecordSRV{}, priErr
	}

	return RecordSRV{
		Type:     "SRV",
		Name:     name,
		Content:  content,
		TTL:      ttl,
		Priority: priority,
		Weight:   weight,
		Port:     port,
	}, nil
}

// RecordCAA represents Njalla's CAA record
type RecordCAA struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

// GetURLValues converts struct fields back into provider suitable values
func (r RecordCAA) GetURLValues() url.Values {
	values := url.Values{}
	values.Set("id", fmt.Sprintf("%d", r.ID))
	values.Set("type", r.Type)
	values.Set("name", r.Name)
	values.Set("content", r.Content)
	values.Set("ttl", fmt.Sprintf("%d", r.TTL))
	return values
}

// GetID exposes the internal ID
func (r RecordCAA) GetID() int {
	return r.ID
}

// NewRecordCAA validates and creates a new RecordCAA
func NewRecordCAA(name string, content string, ttl int) (RecordCAA, error) {
	if ttlErr := checkValidTTL(ttl); ttlErr != nil {
		return RecordCAA{}, ttlErr
	}

	return RecordCAA{
		Type:    "CAA",
		Name:    name,
		Content: content,
		TTL:     ttl,
	}, nil
}

// RecordPTR represents Njalla's PTR record
type RecordPTR struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

// GetURLValues converts struct fields back into provider suitable values
func (r RecordPTR) GetURLValues() url.Values {
	values := url.Values{}
	values.Set("id", fmt.Sprintf("%d", r.ID))
	values.Set("type", r.Type)
	values.Set("name", r.Name)
	values.Set("content", r.Content)
	values.Set("ttl", fmt.Sprintf("%d", r.TTL))
	return values
}

// GetID exposes the internal ID
func (r RecordPTR) GetID() int {
	return r.ID
}

// NewRecordPTR validates and creates a new RecordPTR
func NewRecordPTR(name string, content string, ttl int) (RecordPTR, error) {
	if ttlErr := checkValidTTL(ttl); ttlErr != nil {
		return RecordPTR{}, ttlErr
	}

	return RecordPTR{
		Type:    "PTR",
		Name:    name,
		Content: content,
		TTL:     ttl,
	}, nil
}

// RecordNS represents Njalla's NS record
type RecordNS struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

// GetURLValues converts struct fields back into provider suitable values
func (r RecordNS) GetURLValues() url.Values {
	values := url.Values{}
	values.Set("id", fmt.Sprintf("%d", r.ID))
	values.Set("type", r.Type)
	values.Set("name", r.Name)
	values.Set("content", r.Content)
	values.Set("ttl", fmt.Sprintf("%d", r.TTL))
	return values
}

// GetID exposes the internal ID
func (r RecordNS) GetID() int {
	return r.ID
}

// NewRecordNS validates and creates a new RecordNS
func NewRecordNS(name string, content string, ttl int) (RecordNS, error) {
	if ttlErr := checkValidTTL(ttl); ttlErr != nil {
		return RecordNS{}, ttlErr
	}

	return RecordNS{
		Type:    "NS",
		Name:    name,
		Content: content,
		TTL:     ttl,
	}, nil
}

// RecordTLSA represents Njalla's TLSA record
type RecordTLSA struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

// GetURLValues converts struct fields back into provider suitable values
func (r RecordTLSA) GetURLValues() url.Values {
	values := url.Values{}
	values.Set("id", fmt.Sprintf("%d", r.ID))
	values.Set("type", r.Type)
	values.Set("name", r.Name)
	values.Set("content", r.Content)
	values.Set("ttl", fmt.Sprintf("%d", r.TTL))
	return values
}

// GetID exposes the internal ID
func (r RecordTLSA) GetID() int {
	return r.ID
}

// NewRecordTLSA validates and creates a new RecordTLSA
func NewRecordTLSA(name string, content string, ttl int) (RecordTLSA, error) {
	if ttlErr := checkValidTTL(ttl); ttlErr != nil {
		return RecordTLSA{}, ttlErr
	}

	return RecordTLSA{
		Type:    "TLSA",
		Name:    name,
		Content: content,
		TTL:     ttl,
	}, nil
}

// RecordRedirect represents Njalla's Redirect record
type RecordRedirect struct {
	ID           int    `json:"id"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	URL          string `json:"content"`
	RedirectType int    `json:"prio"`
}

// GetURLValues converts struct fields back into provider suitable values
func (r RecordRedirect) GetURLValues() url.Values {
	values := url.Values{}
	values.Set("id", fmt.Sprintf("%d", r.ID))
	values.Set("type", r.Type)
	values.Set("name", r.Name)
	values.Set("content", r.URL)
	values.Set("prio", fmt.Sprintf("%d", r.RedirectType))
	return values
}

// GetID exposes the internal ID
func (r RecordRedirect) GetID() int {
	return r.ID
}

// NewRecordRedirect validates and creates a new RecordRedirect
func NewRecordRedirect(name string, url string, redirectType int) (RecordRedirect, error) {
	if rtypeErr := checkValidPriority(redirectType); rtypeErr != nil {
		return RecordRedirect{}, rtypeErr
	}

	return RecordRedirect{
		Type:         "Redirect",
		Name:         name,
		URL:          url,
		RedirectType: redirectType,
	}, nil
}

// RecordDynamic represents Njalla's Dynamic record
type RecordDynamic struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	TTL  int    `json:"ttl"`
}

// GetURLValues converts struct fields back into provider suitable values
func (r RecordDynamic) GetURLValues() url.Values {
	values := url.Values{}
	values.Set("id", fmt.Sprintf("%d", r.ID))
	values.Set("type", r.Type)
	values.Set("name", r.Name)
	values.Set("ttl", fmt.Sprintf("%d", r.TTL))
	return values
}

// GetID exposes the internal ID
func (r RecordDynamic) GetID() int {
	return r.ID
}

// NewRecordDynamic validates and creates a new RecordDynamic
func NewRecordDynamic(name string, content string, ttl int) (RecordDynamic, error) {
	if ttlErr := checkValidTTL(ttl); ttlErr != nil {
		return RecordDynamic{}, ttlErr
	}

	return RecordDynamic{
		Type: "Dynamic",
		Name: name,
		TTL:  ttl,
	}, nil
}

// RecordSSHFP represents Njalla's SSHFP record
type RecordSSHFP struct {
	ID           int    `json:"id"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	TTL          int    `json:"ttl"`
	SSHAlgorithm int    `json:"ssh_algorithm"`
	SSHType      int    `json:"ssh_type"`
	Content      string `json:"content"`
}

// GetURLValues converts struct fields back into provider suitable values
func (r RecordSSHFP) GetURLValues() url.Values {
	values := url.Values{}
	values.Set("id", fmt.Sprintf("%d", r.ID))
	values.Set("type", r.Type)
	values.Set("name", r.Name)
	values.Set("content", r.Content)
	values.Set("ttl", fmt.Sprintf("%d", r.TTL))
	values.Set("ssh_algorithm", fmt.Sprintf("%d", r.SSHAlgorithm))
	values.Set("ssh_type", fmt.Sprintf("%d", r.SSHType))
	return values
}

// GetID exposes the internal ID
func (r RecordSSHFP) GetID() int {
	return r.ID
}

// NewRecordSSHFP validates and creates a new RecordSSHFP
func NewRecordSSHFP(
	name string, content string, ttl int, sshAlgorithm int, sshType int,
) (RecordSSHFP, error) {
	if ttlErr := checkValidTTL(ttl); ttlErr != nil {
		return RecordSSHFP{}, ttlErr
	}

	if sshAlgErr := checkValidSSHAlgorithm(sshAlgorithm); sshAlgErr != nil {
		return RecordSSHFP{}, sshAlgErr
	}

	if sshTypeErr := checkValidSSHType(sshType); sshTypeErr != nil {
		return RecordSSHFP{}, sshTypeErr
	}

	return RecordSSHFP{
		Type:         "SSHFP",
		Name:         name,
		TTL:          ttl,
		SSHAlgorithm: sshAlgorithm,
		SSHType:      sshType,
		Content:      content,
	}, nil
}

func checkValidTTL(ttl int) error {
	valid := []int{
		structures.TTL60, structures.TTL300, structures.TTL900,
		structures.TTL10800, structures.TTL21600, structures.TTL86400,
	}

	for _, value := range valid {
		if ttl == value {
			return nil
		}
	}

	return fmt.Errorf("Given TTL [%d] is not valid: %+v", ttl, valid)
}

func checkValidPriority(priority int) error {
	valid := []int{
		structures.PRIORITY0, structures.PRIORITY1, structures.PRIORITY5,
		structures.PRIORITY10, structures.PRIORITY20, structures.PRIORITY30,
		structures.PRIORITY40, structures.PRIORITY50, structures.PRIORITY60,
	}

	for _, value := range valid {
		if priority == value {
			return nil
		}
	}

	return fmt.Errorf(
		"Given Priority [%d] is not valid: %+v", priority, valid,
	)
}

func checkValidRedirectType(redirectType int) error {
	valid := []int{structures.REDIRECTTYPE301, structures.REDIRECTTYPE302}

	for _, value := range valid {
		if redirectType == value {
			return nil
		}
	}

	return fmt.Errorf(
		"Given Redirect Type [%d] is not valid: %+v", redirectType, valid,
	)
}

func checkValidSSHAlgorithm(sshAlgorithm int) error {
	valid := []int{
		structures.SSHALGORITHMRSA, structures.SSHALGORITHMDSA,
		structures.SSHALGORITHMECDSA, structures.SSHALGORITHMED25519,
	}

	for _, value := range valid {
		if sshAlgorithm == value {
			return nil
		}
	}

	return fmt.Errorf(
		"Given SSH Algorithm [%d] is not valid: %+v", sshAlgorithm, valid,
	)
}

func checkValidSSHType(sshType int) error {
	valid := []int{structures.SSHTYPESSHA1, structures.SSHTYPESSHA256}

	for _, value := range valid {
		if sshType == value {
			return nil
		}
	}

	return fmt.Errorf(
		"Given SSH Type [%d] is not valid: %+v", sshType, valid,
	)
}
