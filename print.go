package ansi

import "fmt"

const esc = "\033[" // esc is the initial ANSI escape code (\033 or \x1b)

// Print overrides the fmt.Print function by adding the line and column
// parameters to be able to position a string in a part of the terminal
// following the ANSI standard. The default behavior of fmt.Print is maintained.
func Print(line, col int, str string) {
	setLineCol(line, col)
	fmt.Print(str)
}

// Printf overrides the fmt.Printf function by adding the line and column
// parameters to be able to position a string in a part of the terminal
// following the ANSI standard. The default behavior of fmt.Printf is maintained.
func Printf(line, col int, format string, a ...any) {
	setLineCol(line, col)
	fmt.Printf(format, a...)
}

// setLineCol builds the ansi code to set the cursor position from the line and column values.
func setLineCol(line, col int) {
	fmt.Printf("%s%d;%dH", esc, line, col)
}
