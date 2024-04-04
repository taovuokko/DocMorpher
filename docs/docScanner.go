package docs

import (
	"os"
	"strings"
)

// ListDocxTemplates scans the specified directory for .docx files and returns a slice of their names.
func ListDocxTemplates(dirPath string) ([]string, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	var docxFiles []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue // Skip directories
		}
		if strings.HasSuffix(entry.Name(), ".docx") {
			docxFiles = append(docxFiles, entry.Name())
		}
	}

	if len(docxFiles) == 0 {
		return nil, os.ErrNotExist
	}

	return docxFiles, nil
}
