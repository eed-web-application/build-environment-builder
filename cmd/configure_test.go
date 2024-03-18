package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigure(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"configure"})

	err := rootCmd.Execute()
	assert.NoError(t, err)
}

func TestConfigureAddEndpoint(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"configure", "endpoint", "--label=dev", "--url=http://localhost:8080/api/v1/"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	// check the saved configuration
	assert.Equal(t, "http://localhost:8080/api/v1/", Configuration.Endpoints["dev"])
}

func TestConfigureAndDeleteEndpoint(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"configure", "endpoint", "--label=dev", "--url=http://localhost:8080/api/v1/"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	// check the saved configuration
	assert.Equal(t, "http://localhost:8080/api/v1/", Configuration.Endpoints["dev"])

	rootCmd.SetArgs([]string{"configure", "endpoint", "--label=dev", "--url="})

	err = rootCmd.Execute()
	assert.NoError(t, err)

	// check the saved configuration
	_, ok := Configuration.Endpoints["dev"]
	assert.False(t, ok, "The 'dev' key should be removed from the Endpoints map.")
}
