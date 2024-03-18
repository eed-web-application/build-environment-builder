package cbs_test

import (
	"testing"

	"github.com/eed-web-application/build-environment-builder/cbs"
	"github.com/stretchr/testify/assert"
)

func TestCreateComponent(t *testing.T) {
	id, err := cbs.CreateNewComponent(
		cbs.NewComponentDTO{
			Name:    "test",
			Url:     "component url",
			Version: "1.0.0",
		},
	)
	assert.NoError(t, err)
	assert.NotNil(t, id, "Id should be valorized")
	found_component, err := cbs.FindAllComponent()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(*found_component), "The length of the component array should be 1")
}
