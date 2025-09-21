package errors

import (
	"fmt"
	"os"
)

func HandleFileError(commandName string, fileName string, err error) bool {
	if err == nil {
		return false
	}

	switch {
	case os.IsNotExist(err):
		//file not found
		fmt.Fprintf(os.Stderr, "Error in %s: file not found: %s\n", commandName, fileName)
		return true
	case os.IsPermission(err):
		fmt.Fprintf(os.Stderr, "Error in %s: permission denied to access file: %s\n", commandName, fileName)
		return true
	default:
		// a generic error message for all other file errors
		fmt.Fprintf(os.Stderr, "Error in %s: failed to handle file %s: %v \n", commandName, fileName, err)
		return true
	}

}
