package records

// Record type
type Record struct {
	Info map[string]interface{}
}

// RecordA type
type RecordA struct {
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

// Records interface implemented by every record
type Records interface {
	getValue()
	getInfo()
	setType()
}

// New constructor
func New(infoScraped map[string]interface{}) *Record {
	return &Record{
		Info: infoScraped,
	}
}

func (r *Record) getValue(key string) interface{} {
	return r.Info[key]
}

func (r *Record) getInfo() map[string]interface{} {
	return r.Info
}
