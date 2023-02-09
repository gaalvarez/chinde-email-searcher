package main

import (
	"fmt"
	"galvarezl/mail-indexer/files"
	"galvarezl/mail-indexer/indexer"
	"galvarezl/mail-indexer/parser"
	"sync"
	"time"
)

// create const for the root directory
const root = "/Users/gustavoalvarez/Documents/personal/Truora/enron_mail_20110402/maildir/"

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
			json := parser.JSONToIndexByUser(emails, folder)
			indexer.IndexUserRecords(json, folder)
		}(folder)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("It took %s", elapsed)
}
