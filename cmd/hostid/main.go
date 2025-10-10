package main

/*
 #include <unistd.h>
 #include <stdint.h>
 #include <stdlib.h>

 unsigned long get_hostid() {
     return gethostid();
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
)

const HELP = `Usage: hostid [OPTION]
Print the numeric identifier (in hexadecimal) for the current host.

      --help        display this help and exit
      --version     output version information and exit

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/hostid>
or available locally via: info '(coreutils) hostid invocation'`

var rootCmd = &cobra.Command{
	Use:   "Usage: hostid [OPTION]",
	Short: "Print the numeric identifier (in hexadecimal) for the current host.",
	Long:  HELP,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Println(utils.HOSTID_VERSION)
			return
		}
		id := C.get_hostid()
		fmt.Printf("%08x\n", uint32(id)&0xffffffff)
		return
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "output version information and exit")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
