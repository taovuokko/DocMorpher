package tui

import (
	"docmorph/docs"
	"log"
	"os"

	"github.com/rivo/tview"
)

// UserData holds the input data from the user
type UserData struct {
	OrderNumber      string
	Name             string
	ActivationCode   string
	SelectedTemplate string
}

// StartTUI initializes and displays the Terminal User Interface
func StartTUI() *UserData {
	app := tview.NewApplication().EnableMouse(true)

	// Create a new form.
	form := tview.NewForm().
		AddInputField("Order Number", "", 20, nil, nil).
		AddInputField("Name", "", 20, nil, nil).
		AddInputField("Activation Code", "", 20, nil, nil)

	// Attempt to list .docx templates from the specified directory.
	templateDir := "./templates"
	templateList, err := docs.ListDocxTemplates(templateDir)
	if err != nil {
		log.Fatalf("Failed to list .docx templates: %v", err)
	}

	// Variable to hold the selected template.
	var selectedTemplate string

	// Dropdown for template selection.
	form.AddDropDown("Select Template: ", templateList, 0, func(option string, index int) {
		selectedTemplate = option // Update the selected template based on the user's selection.
	})

	// Submit button to finalize the form.
	form.AddButton("Submit", func() {
		app.Stop()
	})

	// Exit button to exit the application.
	form.AddButton("Exit", func() {
		os.Exit(0)
	})

	// Set the form as the root and focus it.
	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
		log.Fatalf("Failed to start TUI: %v", err)
	}

	// Retrieve the data from the form fields.
	userData := &UserData{
		OrderNumber:      form.GetFormItemByLabel("Order Number").(*tview.InputField).GetText(),
		Name:             form.GetFormItemByLabel("Name").(*tview.InputField).GetText(),
		ActivationCode:   form.GetFormItemByLabel("Activation Code").(*tview.InputField).GetText(),
		SelectedTemplate: selectedTemplate,
	}

	return userData
}
