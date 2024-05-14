package utility

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

// Use the `//go:embed` directive with the path to your template file or directory.
// This tells Go to embed the file(s) in the variable `templateFS`.
//
//go:embed templates/new_component.yaml
//go:embed templates/new_command.yaml
var templateFS embed.FS

// GetTemplate reads the content of a template file from the embedded filesystem.
func GetTemplate(name string) (string, error) {
	templateContent, err := fs.ReadFile(templateFS, fmt.Sprintf("templates/%v", name))
	if err != nil {
		return "", err
	}

	// Create a temporary file for the new resource.
	tempFile, err := os.CreateTemp("", "resource-*.yaml")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	// Write the template content to the temporary file.
	_, err = tempFile.Write(templateContent)
	if err != nil {
		return "", err
	}
	return tempFile.Name(), nil
}
