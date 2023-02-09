package data

type Email struct {
	MessageID string
	Subject   string
	Date      string
	From      string
	To        []string
	Cc        []string
	Bcc       []string
	Body      string
	// MimeVersion  string
	// ContentType  string
	// ContentEncoding string
	// XFrom string
	// XTo   string
	// XCc   string
	// XBcc  string
	// XFolder      string
	// XOrigin      string
	// XFileName    string
}

type IndexerPayload struct {
	Index   string  `json:"index"`
	Records []Email `json:"records"`
}
