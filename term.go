package ansi

import "fmt"

const (
	ansiClear      = "\033[H\033[2J"
	ansiHideCursor = "\033[?25l"
	ansiShowCursor = "\033[?25h"
)

// Clear clears the terminal by removing all printed characters
func Clear() {
	fmt.Print(ansiClear)
}

// HideCursor hides the cursor.
// This option is permanent in the current terminal settings until
// it is destroyed. A new terminal should not have the cursor hidden.
func HideCursor() {
	fmt.Print(ansiHideCursor)
}

// ShowCursor shows the cursor if it is hidden.
// This option is permanent in the current terminal settings until
// it is destroyed.
func ShowCursor() {
	fmt.Print(ansiShowCursor)
}

// Wait is similar to Scanln, stops scanning at a newline and after
// the final item there must be a newline or EOF.
func Wait() {
	fmt.Scanln()
}
