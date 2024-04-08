package cbs

import (
	"bufio"
	"io"
	"os"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// SaveToFile saves the content of the File struct to a physical file.
func WriteToStdOut(f openapi_types.File) error {
	// Write the data to a file.
	writer := bufio.NewWriter(os.Stdout)
	reader, err := f.Reader()
	if err != nil {
		return err
	}
	defer reader.Close()

	// Stream from file to stdout
	if _, err := io.Copy(os.Stdout, reader); err != nil {
		os.Stderr.WriteString("Failed to stream data to stdout\n")
	}

	writer.Flush()
	return nil
}
