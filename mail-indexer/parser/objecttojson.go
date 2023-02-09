package parser

import (
	"encoding/json"
	"fmt"
	"galvarezl/mail-indexer/data"
)

func JSONToIndexByUser(emails []data.Email, username string) string {
	fmt.Println(username, len(emails))
	payload := data.IndexerPayload{
		// Index:   "enronmails",
		Index:   "enronmailsmap1",
		Records: emails,
	}
	return indexByUserToJSON(payload, username)
}

// parse payload to json
func indexByUserToJSON(payload data.IndexerPayload, username string) string {
	bytes, err := json.Marshal(payload)
	if err != nil {
		panic("Is not possible to parse the payload to json for: " + username)
	}
	return string(bytes)
}
