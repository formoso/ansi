package ansi_test

import (
	"bytes"
	"github/formoso/ansi"
	"io"
	"os"
	"testing"
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
	expected := "\x1b[2;23HΑΝΣΙ"
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
	expected := "\x1b[2;23HΑΝΣΙ = ANSI"
	if out != expected {
		t.Errorf("want %v, got %s", expected, out)
	}
}
