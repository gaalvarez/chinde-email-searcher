package main

import (
	"fmt"
	"galvarezl/mail-indexer/files"
	"galvarezl/mail-indexer/indexer"
	"galvarezl/mail-indexer/parser"
	"os"
	"sync"
	"time"
)

// create const for the root directory
var root = os.Getenv("ENRON_ROOT_FOLDER")

func main() {
	start := time.Now()
	folders := files.ReadRootFolders(root)
	var wg sync.WaitGroup
	for _, folder := range folders {
		wg.Add(1)
		go func(folder string) {
			defer wg.Done()
			filePaths := files.FilePaths(root + folder)
			emailTextList := files.ReadFilesContent(filePaths)
			emails := parser.ParseEmails(emailTextList)
			jsons := parser.JSONToIndexByUser(emails, folder)
			for _, json := range jsons {
				indexer.IndexUserRecords(json, folder)
			}
		}(folder)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("It took %s", elapsed)
}
