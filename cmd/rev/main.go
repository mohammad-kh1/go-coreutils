package main

import (
	"bufio"
	"io"
	"os"

	"github.com/mohammad-kh1/go-coreutils/internal/errors"
	"github.com/mohammad-kh1/go-coreutils/internal/utils"

	"github.com/spf13/cobra"
)

var (
	version bool
)

const bufferSize = 1 << 20 // 1MB

const HELP = `Usage:
 rev [options] [<file> ...]

Reverse lines characterwise.

Options:
 -h, --help     display this help
 -V, --version  display version

For more details see rev(1).
`

var rootCmd = &cobra.Command{
	Use:   "rev [options] [<file> ...]",
	Short: "Reverse lines characterwise.",
	Long:  HELP,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			utils.RevVersion()
			return
		}
		if len(args) == 0 {
			//stdin
			reverseString(os.Stdin)
		} else {
			//files
			for _, v := range args {
				//check for file or dir
				fi, errStat := os.Stat(v)
				if errors.HandleFileError("rev", v, errStat) {
					return
				}
				if fi.Mode().IsDir() {
					errors.DirectoryError("rev", v)
					return
				}

				file, err := os.Open(v)
				if errors.HandleFileError("rev", v, err) {
					return
				}
				defer file.Close()
				reverseString(file)
			}
		}

	},
}

func reverseString(r io.Reader) {
	scanner := bufio.NewScanner(r)
	writer := bufio.NewWriterSize(os.Stdout, bufferSize)
	defer writer.Flush()

	for scanner.Scan() {
		input := scanner.Text()
		runes := []rune(input)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		writer.WriteString(string(runes))
		writer.WriteByte('\n')
	}

}

func init() {
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "display version")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
