package mail

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func QueryDataFromZincSearch(searchParams SearchParams) *SearchResults {
	//pass searchParams object to json string

	req, err := http.NewRequest(
		"POST",
		"http://zinc:4080/api/enronmails/_search",
		strings.NewReader(searchObjectToJson(searchParams)),
	)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var searchResponse SearchResponse
	json.Unmarshal(body, &searchResponse)

	return formaData(&searchResponse)
}

func searchObjectToJson(search SearchParams) string {
	bytes, err := json.Marshal(search)
	if err != nil {
		panic("Error converting search object to json")
	}
	return string(bytes)
}

func formaData(sr *SearchResponse) *SearchResults {
	srr := &SearchResults{}
	srr.Data.Total = sr.Hits.Total.Value
	srr.Data.Results = make([]struct {
		Bcc       []interface{} `json:"bcc"`
		Body      string        `json:"body"`
		Cc        []interface{} `json:"cc"`
		Date      string        `json:"date"`
		From      string        `json:"from"`
		MessageID string        `json:"messageID"`
		Subject   string        `json:"subject"`
		To        []string      `json:"to"`
	}, len(sr.Hits.Hits))
	for i, hit := range sr.Hits.Hits {
		srr.Data.Results[i].Bcc = hit.Source.Bcc
		srr.Data.Results[i].Body = hit.Source.Body
		srr.Data.Results[i].Cc = hit.Source.Cc
		srr.Data.Results[i].Date = hit.Source.Date
		srr.Data.Results[i].From = hit.Source.From
		srr.Data.Results[i].MessageID = hit.Source.MessageID
		srr.Data.Results[i].Subject = hit.Source.Subject
		srr.Data.Results[i].To = hit.Source.To
	}
	return srr
}
