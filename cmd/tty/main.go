package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <errno.h>

#define STDIN_FD 0

int tty(void){

	char *tty_name;

	if (isatty(STDIN_FD) == 0) {
		puts("not a tty");
		return EXIT_FAILURE;
	}
	tty_name = ttyname(STDIN_FD);

	if (tty_name == NULL){
		// ttyname failed, even though it's a TTY (e.g., permission error).
        // Print an error message to stderr and exit with status 1.
        fprintf(stderr, "my_tty: Error retrieving TTY name (errno %d)\n", errno);
        return EXIT_FAILURE;
	}

	printf("%s\n", tty_name);
	return 0;

}

*/
import "C"

import (
	"fmt"
	"os"

	"github.com/mohammad-kh1/go-coreutils/internal/utils"
	"github.com/spf13/cobra"
)

var (
	version bool
	silent  bool
)

const HELP = `Usage: tty [OPTION]...
Print the file name of the terminal connected to standard input.

  -s, --silent, --quiet   print nothing, only return an exit status
      --help        display this help and exit
      --version     output version information and exit

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/tty>
or available locally via: info '(coreutils) tty invocation'
`

var rootCmd = &cobra.Command{
	Use:   "Usage: tty [OPTION]...",
	Short: "Print the file name of the terminal connected to standard input.",
	Long:  HELP,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Println(utils.TTY_VERSION)
			return
		}
		if silent {
			return
		}

		C.tty()

	},
}

func init() {
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "output version information and exit")
	rootCmd.Flags().BoolVarP(&silent, "silent", "s", false, "print nothing, only return an exit status")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(10)
	}
}
