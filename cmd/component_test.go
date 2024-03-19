package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComponent(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"component"})

	err := rootCmd.Execute()
	assert.NoError(t, err)
}

func TestComponentFindAllfailNoLabel(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"component", "find-all"})

	err := rootCmd.Execute()
	assert.Error(t, err)
}

func TestComponentFindAllfailLabelNotFound(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"component", "find-all", "--label", "notfound"})

	err := rootCmd.Execute()
	assert.Error(t, err)
}

func TestComponentFindAllOK(t *testing.T) {
	Configuration.Endpoints = make(map[string]string)
	Configuration.Endpoints["test-1"] = "http://cbs:8080"
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"component", "find-all", "--label", "test-1"})

	err := rootCmd.Execute()
	assert.NoError(t, err)
}
