package main

import (
	"docmorph/docs"
	"docmorph/tui"
	"fmt"
	"log"
)

// Converts tui.UserData to docs.UserData
func convertUserData(tuiData *tui.UserData) *docs.UserData {
	return &docs.UserData{
		OrderNumber:      tuiData.OrderNumber,
		Name:             tuiData.Name,
		ActivationCode:   tuiData.ActivationCode,
		SelectedTemplate: tuiData.SelectedTemplate,
	}
}

func main() {
	fmt.Println("DocMorph Application Starting...")

	userData := tui.StartTUI()
	log.Printf("Collected User Data: %+v\n", userData)

	// Convert userData type before calling ModifyTemplate function
	docsUserData := convertUserData(userData)

	templatePath := "./templates/" + userData.SelectedTemplate

	modifiedDocPath, originalTemplateName, err := docs.ModifyTemplate(templatePath, docsUserData)
	if err != nil {
		log.Fatalf("Error modifying template: %v\n", err)
	}
	log.Printf("Template successfully modified: %s\n", modifiedDocPath)

	// Use originalTemplateName when calling ConvertToPDF function
	outputDir := "./output"
	pdfFilePath, err := docs.ConvertToPDF(modifiedDocPath, userData.OrderNumber, originalTemplateName, outputDir)
	if err != nil {
		log.Fatalf("Error converting .docx to PDF: %v\n", err)
	}
	log.Printf("PDF successfully created: %s\n", pdfFilePath)
}
