package ansi_test

import (
	"bytes"
	"fmt"
	"github/formoso/ansi"
	"io"
	"os"
	"testing"
)

func Test_Clear_Success(t *testing.T) {
	var buffer bytes.Buffer
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	ansi.Print(23, 2, "ΑΝΣΙ")
	ansi.Clear()
	w.Close()
	os.Stdout = originalStdout
	io.Copy(&buffer, r)
	out := buffer.String()
	expected := "\x1b[2;23HΑΝΣΙ\x1b[H\x1b[2J"
	if out != expected {
		t.Errorf("want %v, got %s", expected, out)
	}
}

func Test_HideCursor_Success(t *testing.T) {
	var buffer bytes.Buffer
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	ansi.HideCursor()
	w.Close()
	os.Stdout = originalStdout
	io.Copy(&buffer, r)
	out := buffer.String()
	expected := "\x1b[?25l"
	if out != expected {
		t.Errorf("want %v, got %s", expected, out)
	}
}

func Test_ShowCursor_Success(t *testing.T) {
	var buffer bytes.Buffer
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	ansi.ShowCursor()
	w.Close()
	os.Stdout = originalStdout
	io.Copy(&buffer, r)
	out := buffer.String()
	expected := "\033[?25h"
	if out != expected {
		t.Errorf("want %v, got %s", expected, out)
	}
}

func TestWait(t *testing.T) {
	r, w, _ := os.Pipe()
	oriStdin := os.Stdin
	os.Stdin = r
	fmt.Fprintln(w, "ΑΝΣΙ")
	w.Close()
	ansi.Wait()
	os.Stdin = oriStdin
}
