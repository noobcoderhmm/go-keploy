package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test generated using Keploy
func TestMain_PrintHelloGoProject_123(t *testing.T) {
	// Arrange
	// Redirect stdout to capture the output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Act
	main()

	// Close the writer and restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read the captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Assert
	expectedOutput := "Hello, Go project!\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected string")
}

