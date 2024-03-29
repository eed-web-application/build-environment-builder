package cbs_test

import (
	"testing"

	"github.com/eed-web-application/build-environment-builder/cbs"
	"github.com/eed-web-application/build-environment-builder/utility"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewCommandTemplate(t *testing.T) {
	var commandTemplate cbs.NewCommandTemplateDTO
	resetData(GetTestHost())
	// Open the YAML file
	utility.Deserialize("command_template_new_test.yaml", &commandTemplate)

	// Use the parsed data
	id, err := cbs.CreateNewCommandTemplate(GetTestHost(), &commandTemplate)
	assert.NoError(t, err)
	assert.NotNil(t, id, "Id should be valorized")
}

func TestCreateDeleteCommandTemplate(t *testing.T) {
	var commandTemplate cbs.NewCommandTemplateDTO
	resetData(GetTestHost())
	// Open the YAML file
	utility.Deserialize("command_template_new_test.yaml", &commandTemplate)

	// Use the parsed data
	id, err := cbs.CreateNewCommandTemplate(GetTestHost(), &commandTemplate)
	assert.NoError(t, err)
	assert.NotNil(t, id, "Id should be valorized")
	err = cbs.DeleteCommandTemplate(GetTestHost(), *id)
	assert.NoError(t, err)
}

func TestUpdateCommandTemplate(t *testing.T) {
	var commandTemplate cbs.NewCommandTemplateDTO
	var commandTemplateUpdated cbs.UpdateCommandTemplateDTO
	resetData(GetTestHost())
	// Open the YAML file
	utility.Deserialize("command_template_new_test.yaml", &commandTemplate)

	// Use the parsed data
	id, err := cbs.CreateNewCommandTemplate(GetTestHost(), &commandTemplate)
	assert.NoError(t, err)
	assert.NotNil(t, id, "Id should be valorized")
	cmd, err := cbs.FindCommandById(GetTestHost(), *id)
	assert.NoError(t, err)
	assert.NotNil(t, cmd, "command should be valorized")
	// simulate user update
	utility.Deserialize("command_template_update_test.yaml", &commandTemplateUpdated)
	// update
	err = cbs.UpdateCommandById(GetTestHost(), *id, &commandTemplateUpdated)
	assert.NoError(t, err)
	cmd, err = cbs.FindCommandById(GetTestHost(), *id)
	assert.NoError(t, err)
	assert.NotNil(t, cmd, "command should be valorized")
	assert.NotEqual(t, commandTemplateUpdated.Name, cmd.Name, "Name should be updated")
	assert.NotEqual(t, commandTemplateUpdated.Description, cmd.Description, "Name should be updated")
}
