package ui

import "github.com/fatih/color"

// I don't know why I don't wanna use faith here
const ErrPrefix string = "\033[0m[\033[31mERR\033[0m]"

// Some Defined Colors
var (
	Cyan        = color.New(color.FgCyan)
	Yellow      = color.New(color.FgYellow)
	BoldYellow  = color.New(color.FgYellow, color.Bold)
	Blue        = color.New(color.FgBlue)
	Red         = color.New(color.FgRed)
	BoldRed     = color.New(color.FgRed, color.Bold)
	Green       = color.New(color.FgGreen)
	Magenta     = color.New(color.FgMagenta)
	BoldMagenta = color.New(color.FgMagenta, color.Bold)
	Bold        = color.New(color.Bold)
	Reset       = color.New(color.Reset)
)
