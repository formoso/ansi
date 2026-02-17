package ansi_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/formoso/ansi"
)

func Test_Print_Success(t *testing.T) {
	var buffer bytes.Buffer
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	ansi.Print(23, 2, "ΑΝΣΙ")
	w.Close()
	os.Stdout = originalStdout
	io.Copy(&buffer, r)
	out := buffer.String()
	expected := "\x1b[23;2HΑΝΣΙ"
	if out != expected {
		t.Errorf("want %v, got %s", expected, out)
	}
}

func Test_Printf_Success(t *testing.T) {
	var buffer bytes.Buffer
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	ansi.Printf(23, 2, "ΑΝΣΙ = %v", "ANSI")
	w.Close()
	os.Stdout = originalStdout
	io.Copy(&buffer, r)
	out := buffer.String()
	expected := "\x1b[23;2HΑΝΣΙ = ANSI"
	if out != expected {
		t.Errorf("want %v, got %s", expected, out)
	}
}
