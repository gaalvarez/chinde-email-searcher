package main

import (
	"fmt"
	"galvarezl/mail-indexer/files"
	"galvarezl/mail-indexer/indexer"
	"galvarezl/mail-indexer/parser"
	"sync"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
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
