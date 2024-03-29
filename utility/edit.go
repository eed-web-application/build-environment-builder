package utility

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// run an exeternal editor that will edit the file
func EditFile(f string) error {
	// Determine which editor to use
	editor := os.Getenv("EDITOR")
	if editor == "" {
		// Default to 'nano' if EDITOR env var is not set
		editor = "nano"
	}

	// Open the YAML file in the chosen editor
	cmd := exec.Command(editor, filepath.Clean(f))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	logrus.Printf("Opening %s in %s...\n", f, editor)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to open editor: %v", err)
	}
	logrus.Printf("Editor closed. Proceeding with processing the YAML file...")
	return nil
}

// CreateTempFile creates a temporary file with the given YAML content
func CreateTempFile(yaml []byte) (string, error) {
	// Create a temporary file for the new resource
	tempFile, err := os.CreateTemp("", "resource-*.yaml")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	// Write the YAML content to the temporary file
	_, err = tempFile.Write([]byte(yaml))
	if err != nil {
		return "", err
	}
	return tempFile.Name(), nil
}

// Deserialize reads a YAML file and unmarshals it into a Go struct
func Deserialize(file string, object any) error {
	var intermediateData map[string]interface{}
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &intermediateData)
	if err != nil {
		return err
	}
	// Step 2: Convert the intermediate map to JSON
	jsonData, err := json.Marshal(intermediateData)
	if err != nil {
		return err
	}

	// Step 3: Unmarshal JSON into the final Go struct
	if err := json.Unmarshal(jsonData, object); err != nil {
		return err
	}
	return nil
}

// GetColorizedYaml takes a YAML string and returns a colorized version of it
func GetColorizedYaml(yaml string) (*string, error) {
	// Get the YAML lexer
	lexer := lexers.Get("yaml")

	// Get a formatter for terminal output
	formatter := formatters.Get("terminal256")
	if formatter == nil {
		return nil, fmt.Errorf("terminal formatter not found")
	}

	// Use the "friendly" style for coloring
	style := styles.Get("friendly")
	if style == nil {
		style = styles.Fallback
	}

	// Tokenize the YAML content
	iterator, err := lexer.Tokenise(nil, yaml)
	if err != nil {
		return nil, fmt.Errorf("trror tokenizing YAML content: %v", err)
	}

	// Create a buffer to hold the formatted output
	var b bytes.Buffer

	// Format the tokens into the buffer
	err = formatter.Format(&b, style, iterator)
	if err != nil {
		return nil, fmt.Errorf("error formatting YAML content: %v", err)
	}

	// Print the colorized YAML
	res := b.String()
	return &res, nil
}

func GetFileModTime(filePath string) (time.Time, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return time.Time{}, err
	}
	return fileInfo.ModTime(), nil
}
