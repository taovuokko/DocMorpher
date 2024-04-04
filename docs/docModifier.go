package docs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lukasjarosch/go-docx"
)

type UserData struct {
	OrderNumber      string
	Name             string
	ActivationCode   string
	SelectedTemplate string // Tämä oletetaan olevan tiedoston nimi ilman polkua
}

// ModifyTemplate modifies a Word document template by replacing placeholders with user data.
// It saves the modified template to a directory and returns the path of the modified file along with the original template name.
func ModifyTemplate(templatePath string, userData *UserData) (string, string, error) {
	doc, err := docx.Open(templatePath)
	if err != nil {
		log.Printf("Failed to open docx file: %s, error: %v", templatePath, err)
		return "", "", err
	}
	defer doc.Close()

	placeholderMap := docx.PlaceholderMap{
		"Name":           userData.Name,
		"ActivationCode": userData.ActivationCode,
	}

	err = doc.ReplaceAll(placeholderMap)
	if err != nil {
		log.Printf("Failed to replace placeholders in docx file: %s, error: %v", templatePath, err)
		return "", "", err
	}

	modifiedDir := "./modified"
	if _, err := os.Stat(modifiedDir); os.IsNotExist(err) {
		err := os.MkdirAll(modifiedDir, os.ModePerm)
		if err != nil {
			log.Printf("Failed to create modified directory: %v", err)
			return "", "", err
		}
	}

	// Generate a unique filename for the modified template
	newFilePath := filepath.Join(modifiedDir, fmt.Sprintf("%s_%s_modified.docx", userData.OrderNumber, time.Now().Format("20060102_150405")))
	err = doc.WriteToFile(newFilePath)
	if err != nil {
		log.Printf("Failed to save modified docx file: %s, error: %v", newFilePath, err)
		return "", "", err
	}

	log.Printf("Template successfully modified and saved: %s", newFilePath)

	// Return the path of the modified file and the original template name (without path and extension)
	originalTemplateName := filepath.Base(userData.SelectedTemplate)
	originalTemplateName = strings.TrimSuffix(originalTemplateName, filepath.Ext(originalTemplateName))
	return newFilePath, originalTemplateName, nil
}
