package parser

import (
	"encoding/json"
	"fmt"
	"galvarezl/mail-indexer/data"
)

// func JSONToIndexByUser(emails []data.Email, username string) string {
// 	fmt.Println(username, len(emails))
// 	payload := data.IndexerPayload{
// 		// Index: "enronmails",
// 		Index:   "enronmailsmap1",
// 		Records: emails,
// 	}
// 	return indexByUserToJSON(payload, username)
// }

func JSONToIndexByUser(emails []data.Email, username string) []string {
	fmt.Println(username, len(emails))
	const indexName = "enronmails"
	// const indexName = "enronmailsmap2"
	if len(emails) > 7000 {
		// Dividir el payload en dos partes
		halfLength := len(emails) / 2
		payload1 := data.IndexerPayload{
			Index:   indexName,
			Records: emails[:halfLength],
		}
		payload2 := data.IndexerPayload{
			Index:   indexName,
			Records: emails[halfLength:],
		}
		// Llamar a la función indexByUserToJSON dos veces
		return []string{indexByUserToJSON(payload1, username), indexByUserToJSON(payload2, username)}
	} else {
		payload := data.IndexerPayload{
			// Index: "enronmails",
			Index:   indexName,
			Records: emails,
		}
		return []string{indexByUserToJSON(payload, username)}
	}
}

// parse payload to json
func indexByUserToJSON(payload data.IndexerPayload, username string) string {
	bytes, err := json.Marshal(payload)
	if err != nil {
		panic("Is not possible to parse the payload to json for: " + username)
	}
	return string(bytes)
}
