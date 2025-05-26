# ΑΝΣΙ
ΑΝΣΙ is a library that helps work with the ANSI standard in terminals using Golang.

<img src="/assets/go-gopher-ansi.png" width="200" alt="ANSI Gopher" />


## What is the ANSI Standard?
ANSI (American National Standards Institute) is an organization that defines standards across various technologies — including control codes for text terminals.

The ANSI standard for terminals defines a set of escape codes that allow you to control the appearance and behavior of the terminal using text-based commands.

These codes are used to:
- Move the cursor to a specific position
- Clear the screen or a line
- Change text and background colors
- Style text (bold, underline, etc.)
- Hide or show the cursor
- Create simple character-based animations

## How Do ANSI Codes Work?

ANSI escape codes begin with an **escape character** (`\033` or `\x1b`), followed by a bracket and a command:

### Example:
```go
fmt.Print("\033[31mHello\033[0m")
```

## Main Features
- Position a character or a string at a specific point in the terminal.
- Provide basic functions such as:
    - Clearing the terminal;
    - Hiding the cursor;
    - Revealing the cursor.

## Features to be added in future versions
- Change the color of a character and the terminal;
- Move a character or a string across the terminal;
- Provide basic functions such as:
    - Getting the terminal area size;
    - Setting the terminal area size.

## Example of how to use
```go
package main

import "github/formoso/ansi"

func main() {
	// Clears the terminal
	ansi.Clear()

	// Hides the cursor
	ansi.HideCursor()

	// Prints "ΑΝΣΙ" at line 23, column 2
	ansi.Print(23, 2, "ΑΝΣΙ")

	// Expects to receive the escape code \n
	ansi.Wait()

	// Shows the cursor
	ansi.ShowCursor()
}
```