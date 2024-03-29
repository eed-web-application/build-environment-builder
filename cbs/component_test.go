package cbs_test

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/eed-web-application/build-environment-builder/cbs"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestCreateNewCommandTemplate(t *testing.T) {
	// Open the YAML file
	filePath := filepath.Join("command_tempalte_test.yaml")
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	// Parse the YAML content
	var intermediateData map[string]interface{}
	err = yaml.Unmarshal(yamlFile, &intermediateData)
	if err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}
	// Step 2: Convert the intermediate map to JSON
	jsonData, err := json.Marshal(intermediateData)
	if err != nil {
		log.Fatalf("Failed to marshal intermediate data to JSON: %v", err)
	}

	// Step 3: Unmarshal JSON into the final Go struct
	var commandTemplate cbs.NewCommandTemplateDTO
	if err := json.Unmarshal(jsonData, &commandTemplate); err != nil {
		log.Fatalf("Failed to unmarshal JSON into Go struct: %v", err)
	}

	// Use the parsed data
	id, err := cbs.CreateNewCommandTemplate("http://cbs:8080", commandTemplate)
	assert.NoError(t, err)
	assert.NotNil(t, id, "Id should be valorized")

	// delete command
	cbs.DeleteCommandTemplate("http://cbs:8080", *id)
	assert.NoError(t, err)
}

func TestCreateComponent(t *testing.T) {
	id, err := cbs.CreateNewComponent(
		"http://cbs:8080",
		cbs.NewComponentDTO{
			Name:    "test",
			Version: "1.0.0",
		},
	)
	assert.NoError(t, err)
	assert.NotNil(t, id, "Id should be valorized")
	found_component, err := cbs.FindAllComponent("http://cbs:8080")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(*found_component), "The length of the component array should be 1")
}
