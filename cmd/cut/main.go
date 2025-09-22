package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mohammad-kh1/go-coreutils/internal/errors"
	"github.com/spf13/cobra"
)

var (
	delimiter string
	field     int
)

const HELP = `
Usage: cut OPTION... [FILE]...
Print selected parts of lines from each FILE to standard output.

With no FILE, or when FILE is -, read standard input.

Mandatory arguments to long options are mandatory for short options too.
  -b, --bytes=LIST        select only these bytes
  -c, --characters=LIST   select only these characters
  -d, --delimiter=DELIM   use DELIM instead of TAB for field delimiter
  -f, --fields=LIST       select only these fields;  also print any line
                            that contains no delimiter character, unless
                            the -s option is specified
  -n                      (ignored)
      --complement        complement the set of selected bytes, characters
                            or fields
  -s, --only-delimited    do not print lines not containing delimiters
      --output-delimiter=STRING  use STRING as the output delimiter
                            the default is to use the input delimiter
  -z, --zero-terminated   line delimiter is NUL, not newline
      --help        display this help and exit
      --version     output version information and exit

Use one, and only one of -b, -c or -f.  Each LIST is made up of one
range, or many ranges separated by commas.  Selected input is written
in the same order that it is read, and is written exactly once.
Each range is one of:

  N     N'th byte, character or field, counted from 1
  N-    from N'th byte, character or field, to end of line
  N-M   from N'th to M'th (included) byte, character or field
  -M    from first to M'th (included) byte, character or field

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/cut>
or available locally via: info '(coreutils) cut invocation'
`

var rootCmd = &cobra.Command{
	Use:   "Usage: cut OPTION... [FILE]...",
	Short: "Print selected parts of lines from each FILE to standard output.",
	Long:  HELP,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if field == 0 {
			errors.CutFieldError("cut")
			return
		}
		if len(args) == 0 {
			//stdin
			cutPrint(os.Stdin)
			return
		}
		for _, v := range args {
			//check for file or dir
			fi , errStat := os.Stat(v)
			if errors.HandleFileError("cut" , v , errStat){
				return
			}
			if fi.Mode().IsDir(){
				errors.DirectoryError("cut" , v)
				return
			}

			//reading from files
			file, err := os.Open(v)
			
			if errors.HandleFileError("cut", v, err) {
				return
			}
			cutPrint(file)
			defer file.Close()
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&delimiter, "delimiter", "d", " ", "use DELIM instead of TAB for field delimiter")
	rootCmd.Flags().IntVarP(&field, "fields", "f", 0, "select only these fields;  also print any line")
}

func cutPrint(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), delimiter)
		if field > 0 && field <= len(fields) {
			fmt.Println(fields[field-1])
		}
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
