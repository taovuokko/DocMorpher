package tui

import (
	"docmorph/docs"
	"log"

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
	app := tview.NewApplication()

	// Create a new form.
	form := tview.NewForm().
		AddInputField("Order Number", "", 20, nil, func(text string) {
			// Input validation or manipulation can be done here.
		}).
		AddInputField("Name", "", 20, nil, nil).
		AddInputField("Activation Code", "", 20, nil, nil)
	form.SetBorder(true).SetTitle("Enter Details").SetTitleAlign(tview.AlignLeft)

	// Attempt to list .docx templates from the specified directory.
	templateDir := "./templates"
	TemplateList, err := docs.ListDocxTemplates(templateDir)
	if err != nil {
		log.Fatalf("Failed to list .docx templates: %v", err)
	}

	// Dropdown for template selection.
	var selectedTemplate string
	form.AddDropDown("Select Template: ", TemplateList, 0, func(option string, index int) {
		selectedTemplate = option
		log.Printf("Template selected: %s", selectedTemplate) // Logging the selected template
	})

	// Submit button to finalize the form.
	form.AddButton("Submit", func() {
		app.Stop()
	})
	form.SetBorder(true).SetTitle("Enter Details").SetTitleAlign(tview.AlignLeft)

	if err := app.SetRoot(form, true).Run(); err != nil {
		log.Fatalf("Failed to start TUI: %v", err)
	}

	// Retrieve the data from the form fields.
	userData := &UserData{
		OrderNumber:      form.GetFormItemByLabel("Order Number").(*tview.InputField).GetText(),
		Name:             form.GetFormItemByLabel("Name").(*tview.InputField).GetText(),
		ActivationCode:   form.GetFormItemByLabel("Activation Code").(*tview.InputField).GetText(),
		SelectedTemplate: selectedTemplate,
	}

	// Log collected user data.
	log.Printf("User Data Collected: OrderNumber: %s, Name: %s, ActivationCode: %s, SelectedTemplate: %s",
		userData.OrderNumber, userData.Name, userData.ActivationCode, userData.SelectedTemplate)

	return userData
}
