# DocMorph

DocMorph is a simple Go tool that helps you quickly turn .docx templates into personalized PDFs, ideal for sending out software license details like Microsoft 365 to your customers. 

## Features

- **User-Friendly Interface:** A Terminal User Interface (TUI) simplifies the customization process, requiring only minimal user input.
- **Template Customization:** Users can easily replace specific placeholders (`{Name}` and `{ActivationCode}`) in document templates.
- **PDF Conversion:** Customized templates are converted into PDF format, maintaining the original template's formatting.
- **Document Management:** The application supports listing available .docx templates and saving the final PDFs in a designated output directory.

## Getting Started

### Requirements

- Go 1.22.0 or higher
- LibreOffice: DocMorph uses LibreOffice for the magic part where .docx turns into PDF. It needs to be installed and accessible from your command line or terminal. Check out the LibreOffice website for installation instructions
- Access to a terminal on Windows or Linux platforms

### Quickstart

1. Clone the repository to your local machine.
2. Navigate to the cloned directory.
3. Run `go build` to compile the application.
4. Start the application using `./docmorph` (Linux) or `docmorph.exe` (Windows).
5. Follow the on-screen instructions in the TUI to generate your personalized PDF.
