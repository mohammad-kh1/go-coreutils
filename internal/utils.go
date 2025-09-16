package internal

import (
	"fmt"
	"os"
)

const (
	ErrGeneral = 1 //General errors
	ErrNoFile  = 2 //No such file of directory
)

// Prints a standardized error message to stderr
func PrintError(cmd string, msg string) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", cmd, msg)
}

func FileNotFoundError(cmd string, filename string) int {
	PrintError(cmd, fmt.Sprintf("cannot open '%s': No such file or directory", filename))
	return ErrNoFile
}
