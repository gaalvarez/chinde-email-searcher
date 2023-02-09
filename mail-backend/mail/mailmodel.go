package mail

import "time"

type SearchParams struct {
	SearchType string `json:"search_type"`
	Query      struct {
		Term string `json:"term"`
	} `json:"query"`
	From       int      `json:"from"`
	MaxResults int      `json:"max_results"`
	Source     []string `json:"_source"`
}

type SearchResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index     string    `json:"_index"`
			Type      string    `json:"_type"`
			ID        string    `json:"_id"`
			Score     float64   `json:"_score"`
			Timestamp time.Time `json:"@timestamp"`
			Source    struct {
				Bcc       []interface{} `json:"bcc"`
				Body      string        `json:"body"`
				Cc        []interface{} `json:"cc"`
				Date      string        `json:"date"`
				From      string        `json:"from"`
				MessageID string        `json:"messageID"`
				Subject   string        `json:"subject"`
				To        []string      `json:"to"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type SearchResults struct {
	Data struct {
		Total   int `json:"total"`
		Results []struct {
			Bcc       []interface{} `json:"bcc"`
			Body      string        `json:"body"`
			Cc        []interface{} `json:"cc"`
			Date      string        `json:"date"`
			From      string        `json:"from"`
			MessageID string        `json:"messageID"`
			Subject   string        `json:"subject"`
			To        []string      `json:"to"`
		} `json:"results"`
	} `json:"data"`
}
