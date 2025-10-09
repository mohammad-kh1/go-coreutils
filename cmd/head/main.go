package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/mohammad-kh1/go-coreutils/internal/errors"
	"github.com/mohammad-kh1/go-coreutils/internal/helpers"
	"github.com/mohammad-kh1/go-coreutils/internal/utils"
)

var (
	lines           int
	version         bool
	verbose         bool
	quiet           bool
	zeroTerminated  bool
	bytes           int
)

const HELP = `Usage: head [OPTION]... [FILE]...
Print the first 10 lines of each FILE to standard output.
With more than one FILE, precede each with a header giving the file name.

With no FILE, or when FILE is -, read standard input.

Mandatory arguments to long options are mandatory for short options too.
 -c, --bytes=[-]NUM          print the first NUM bytes of each file;
                             with the leading '-', print all but the last
                             NUM bytes of each file
 -n, --lines=[-]NUM          print the first NUM lines instead of the first 10;
                             with the leading '-', print all but the last
                             NUM lines of each file
 -q, --quiet, --silent       never print headers giving file names
 -v, --verbose               always print headers giving file names
 -z, --zero-terminated       line delimiter is NUL, not newline
     --help                  display this help and exit
     --version               output version information and exit
`

var rootCmd = &cobra.Command{
	Use:   "head [OPTION]... [FILE]...",
	Short: "Print the first 10 lines of each FILE to standard output.",
	Long:  HELP,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Println(utils.HEAD_VERSION)
			return
		}

		// No input files: read from stdin
		if len(args) == 0 {
			// Calculate buffer size for stdin
			bufferSize := helpers.SetBufferSize("-") 
			runHead(os.Stdin, "standard input", true, bufferSize)
			return
		}

		// Multiple input files
		multiple := len(args) > 1
		for _, input := range args {
			runHeadFromFile(input, multiple)
		}
	},
}

// --- main reading logic ---
func runHeadFromFile(filename string, multiple bool) {
	// Calculate buffer size once at the start
	bufferSize := helpers.SetBufferSize(filename)

	if filename == "-" {
		runHead(os.Stdin, "standard input", multiple, bufferSize)
		return
	}

	fi, errStat := os.Stat(filename)
	if errors.HandleFileError("head", filename, errStat) {
		return
	}
	if fi.Mode().IsDir() {
		errors.DirectoryError("head", filename)
		return
	}

	file, err := os.Open(filename)
	if errors.HandleFileError("head", filename, err) {
		return
	}
	runHead(file, filename, multiple, bufferSize)
	defer file.Close()
}

// --- unified print logic ---
func runHead(r io.Reader, label string, multiple bool, bufferSize int) {
	if verbose || (multiple && !quiet) {
		fmt.Printf("==> %s <==\n", label)
	}

	switch {
	case bytes > 0:
		printBytes(r, bytes)
	default:
		printLines(r, bufferSize, lines, zeroTerminated)
	}
}

// --- print by lines (supports -z and proper EOF handling) ---
func printLines(r io.Reader, bufferSize, maxLines int, zeroTerminated bool) {
	reader := bufio.NewReaderSize(r, bufferSize)

	var delim byte
	if zeroTerminated {
		delim = '\x00'
	} else {
		delim = '\n'
	}

	if maxLines <= 0 {
		return
	}

	count := 0
	for {
		data, err := reader.ReadString(delim)

		if len(data) > 0 {
			fmt.Print(data)
			count++
			if count >= maxLines {
				break
			}
		}

		if err == io.EOF {
			break // reached end of input
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "read error:", err)
			break
		}
	}
}

// --- print by bytes (for -c option) ---
func printBytes(r io.Reader, n int) {
	if n <= 0 {
		return
	}
	limited := io.LimitReader(r, int64(n))
	if _, err := io.Copy(os.Stdout, limited); err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, "error copying bytes:", err)
	}
}

// --- init CLI flags ---
func init() {
	rootCmd.Flags().IntVarP(&lines, "lines", "n", 10, "print the first NUM lines instead of the first 10")
	rootCmd.Flags().BoolVarP(&version, "version", "V", false, "output version information and exit")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "always print headers giving file names")
	rootCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "never print headers giving file names")
	rootCmd.Flags().BoolVarP(&zeroTerminated, "zero-terminated", "z", false, "line delimiter is NUL, not newline")
	rootCmd.Flags().IntVarP(&bytes, "bytes", "c", 0, "print the first NUM bytes of each file")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}