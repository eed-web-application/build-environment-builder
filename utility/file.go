package utility

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// WriteToStdOut saves the content of the File struct to a physical file.
func WriteToStdOut(f *bytes.Reader) error {
	// Write the data to a file.
	writer := bufio.NewWriter(os.Stdout)

	// Stream from file to stdout
	if _, err := io.Copy(os.Stdout, f); err != nil {
		os.Stderr.WriteString("Failed to stream data to stdout\n")
	}

	writer.Flush()
	return nil
}
