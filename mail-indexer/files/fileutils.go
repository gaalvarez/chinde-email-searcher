package files

import (
	"os"
	"path/filepath"
	"strings"
)

// List of ignored files const with a element ".DS_Store"
var ignoredFiles = []string{".DS_Store"}

// Read all root folders inside a path
func ReadRootFolders(rootPath string) []string {
	var folders []string
	elements, err := os.ReadDir(rootPath)
	if err != nil {
		panic(err)
	}
	for _, file := range elements {
		if file.IsDir() {
			folders = append(folders, file.Name())
		}
	}
	return folders
}

// read all files inside a root directory and its subdirectories
func FilePaths(folderContainer string) []string {
	var files []string
	err := filepath.Walk(folderContainer, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && !isIgnoredFile(path) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func isIgnoredFile(path string) bool {
	for _, ignoredFile := range ignoredFiles {
		if strings.HasSuffix(path, ignoredFile) {
			return true
		}
	}
	return false
}

// receive a list of files paths in string and for each file path get the content in string format
func ReadFilesContent(files []string) []string {
	var emails []string
	for _, file := range files {
		//check if file is different than Email{}
		emails = append(emails, readContent(file))
	}
	return emails
}

func readContent(file string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(data)
}
