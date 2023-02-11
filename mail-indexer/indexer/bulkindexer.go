package indexer

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func IndexUserRecords(payload string, username string) {
	fmt.Println("start to index user record", username)
	start := time.Now()
	req, err := http.NewRequest("POST", os.Getenv("AWS_PUBLIC_DNS")+":4080/api/_bulkv2", strings.NewReader(payload))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(os.Getenv("ZS_USER"), os.Getenv("ZS_PASSWORD"))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

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
	fmt.Println(string(body))
	elapsed := time.Since(start)
	fmt.Printf("Time for %s, %s", username, elapsed)
	fmt.Println("___________________________________________________________")
}
