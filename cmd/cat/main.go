package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/mohammad-kh1/go-coreutils/internal/errors"
	"github.com/spf13/cobra"
)

var (
	numberLines bool
	numberNonBlank bool
)

const bufferSize = 1 << 20 // 1MB

const HELP = `
With no FILE, or when FILE is -, read standard input.

 -A, --show-all           equivalent to -vET
 -b, --number-nonblank    number nonempty output lines, overrides -n
 -e                       equivalent to -vE
 -E, --show-ends          display $ at end of each line
 -n, --number             number all output lines
 -s, --squeeze-blank      suppress repeated empty output lines
 -t                       equivalent to -vT
 -T, --show-tabs          display TAB characters as ^I
 -u                       (ignored)
 -v, --show-nonprinting   use ^ and M- notation, except for LFD and TAB
     --help        display this help and exit
     --version     output version information and exit

Examples:
   cat f - g  Output f's contents, then standard input, then g's contents.
   cat        Copy standard input to standard output.

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/cat>
or available locally via: info '(coreutils) cat invocation'
`

var rootCmd = &cobra.Command{
	Use:   "Usage: cat [OPTION]... [FILE]...",
	Short: "Concatenate FILE(s) to standard output.",
	Long:  HELP,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			//stdin
			printFile(os.Stdin)
			return
		}
		for _, v := range args {
			if v == "-" {
				printFile(os.Stdin)
				return
			}
			fi , errStat := os.Stat(v)
			if errors.HandleFileError("cat", v, errStat) {
				return 
			}
			if fi.Mode().IsDir(){
				errors.DirectoryError("cat" , v)
			}
			file, err := os.Open(v)
			if errors.HandleFileError("cat", v, err) {
				return 
			}

			printFile(file)

			defer file.Close()
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&numberLines, "number", "n", false, "number all output lines")
	rootCmd.Flags().BoolVarP(&numberNonBlank , "nomber-nonblank" , "b" , false , "number nonempty output lines, overrides -n")
}



func printFile(r io.Reader) {
	scanner := bufio.NewScanner(r)
	writer  := bufio.NewWriterSize(os.Stdout , bufferSize)
	defer writer.Flush()

	lineCount := 0
	for scanner.Scan(){
		text := scanner.Text()


		if numberNonBlank {
			// write line number + tab + text + newline
			if strings.TrimSpace(text) != ""{
				lineCount++
				writer.WriteString("    ")
				writer.WriteString(strconv.Itoa(lineCount))
				writer.WriteString("  ")
				writer.WriteString(text)
				writer.WriteByte('\n')
			}else{
				writer.WriteString("    ")
				writer.WriteString(text)
				writer.WriteByte('\n')
			}

		}else if numberLines{
			lineCount++
			writer.WriteString("    ")
			writer.WriteString(strconv.Itoa(lineCount))
			writer.WriteString("  ")
			writer.WriteString(text)
			writer.WriteByte('\n')

		}else{
			writer.WriteString(text)
			writer.WriteByte('\n')
		}
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
