package cbs_test

import (
	"testing"

	"github.com/eed-web-application/build-environment-builder/cbs"
	"github.com/stretchr/testify/assert"
)

func TestFetchEngines(t *testing.T) {
	// Use the parsed data
	engine_list, err := cbs.FetchAllEngines(GetTestHost())
	assert.NoError(t, err)
	assert.NotNil(t, engine_list, "Egine list should be valorized")
	assert.Equal(t, len(*engine_list), 2, "Egine list should containes two engines")
}
