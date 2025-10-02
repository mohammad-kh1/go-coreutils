package main

import (
	"fmt"
	"runtime"
	"os"
	
	"github.com/mohammad-kh1/go-coreutils/internal/utils"
	"github.com/spf13/cobra"

)

var (
	version bool
)

const HELP = `Usage: arch [OPTION]...
Print machine architecture.

      --help        display this help and exit
      --version     output version information and exit

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/arch>
or available locally via: info '(coreutils) arch invocation'`

var rootCmd = &cobra.Command{
	Use:"Usage: arch [OPTION]...",
	Short:"Print machine architecture.",
	Long:HELP,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command , args []string){
		if version {
			fmt.Println(utils.ARCH_VERSION)
			return
		}

		const goArch = runtime.GOARCH
		archList := map[string]string{
			"amd64":   "x86_64",
			"386":     "i686",
			"arm":     "arm",
			"arm64":   "aarch64",
			"ppc64":   "ppc64",
			"ppc64le": "ppc64le",
			"mips":    "mips",
			"mipsle":  "mipsel",
			"mips64":  "mips64",
			"mips64le":"mips64el",
			"riscv64": "riscv64",
			"s390x":   "s390x",
		}

		if arch , ok := archList[goArch] ; ok {
			fmt.Println(arch)
			return
		}else{
			fmt.Println(goArch)
		}

	},
}

func init(){
	rootCmd.Flags().BoolVarP(&version , "version" , "v" , false , "output version information and exit")
}

func main(){
	if err := rootCmd.Execute() ; err != nil {
		os.Exit(1)
	}
}