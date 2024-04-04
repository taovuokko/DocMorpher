package docs

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// ConvertToPDF converts a modified DOCX file to PDF using the appropriate LibreOffice command based on the operating system.
// It uses the original DOCX file name for the PDF creation and then renames the resulting PDF based on the original template name and order number.
func ConvertToPDF(modifiedDocxPath, orderNumber, originalTemplateName, outputDir string) (string, error) {
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if mkdirErr := os.MkdirAll(outputDir, 0755); mkdirErr != nil {
			log.Printf("Failed to create output directory: %s, error: %v", outputDir, mkdirErr)
			return "", fmt.Errorf("failed to create output directory: %w", mkdirErr)
		}
	}

	// Determine the LibreOffice command path based on the OS
	libreOfficeCmd := "libreoffice" // Default for Linux
	if runtime.GOOS == "windows" {
		libreOfficeCmd = "C:\\Program Files\\LibreOffice\\program\\soffice.exe" // Path for Windows
	} else if runtime.GOOS == "darwin" {
		libreOfficeCmd = "soffice" // Assuming default for MacOS, adjust if necessary
	}

	// Use the original DOCX file name for PDF creation
	cmd := exec.Command(libreOfficeCmd, "--headless", "--convert-to", "pdf", "--outdir", outputDir, modifiedDocxPath)
	if err := cmd.Run(); err != nil {
		log.Printf("Error converting .docx to PDF: %s, error: %v", modifiedDocxPath, err)
		return "", fmt.Errorf("error converting .docx to PDF: %w", err)
	}

	// Expected name of the original PDF file (based on the modified DOCX file name)
	expectedPdfFileName := strings.TrimSuffix(filepath.Base(modifiedDocxPath), filepath.Ext(modifiedDocxPath)) + ".pdf"
	expectedPdfFilePath := filepath.Join(outputDir, expectedPdfFileName)

	// New PDF file name based on the original template name and order number
	newPdfFileName := fmt.Sprintf("Tilaus %s %s.pdf", orderNumber, strings.ReplaceAll(originalTemplateName, ":", " "))
	newPdfFilePath := filepath.Join(outputDir, newPdfFileName)

	// Rename the created PDF file
	if err := os.Rename(expectedPdfFilePath, newPdfFilePath); err != nil {
		log.Printf("Failed to rename PDF file: %v", err)
		return "", err
	}

	log.Printf("PDF successfully created at: %s", newPdfFilePath)
	return newPdfFilePath, nil
}
