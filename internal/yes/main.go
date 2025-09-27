package main

import (
	"os"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

)

var (
	version bool
)



const bufSize = 32 * 1024 // 32KB 
const HELP =`Usage: yes [STRING]... or:  yes OPTION
Repeatedly output a line with all specified STRING(s), or 'y'.

      --help        display this help and exit
      --version     output version information and exit

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/yes>
or available locally via: info '(coreutils) yes invocation'
`


var rootCmd = &cobra.Command{
	Use: "Usage: yes [STRING]... or:  yes OPTION",
	Short: "Repeatedly output a line with all specified STRING(s), or 'y'.",
	Long: HELP,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command , args []string){
		DefaultString:="y\n"

		if version {
			displayVersion()
			return
		}
		
		if len(os.Args) > 1{
			DefaultString = strings.Join(os.Args[1:] , " ") + "\n"
		}
		
		buf := make([]byte , 0 , bufSize)
		for len(buf)+len(DefaultString) <= bufSize{
			buf = append(buf , DefaultString...)
		}

		for {
			if _ , err := os.Stdout.Write(buf) ; err != nil {
				break
			}
		}


	},
}

func init(){
	rootCmd.Flags().BoolVarP(&version , "version" , "n" , false , "output version information and exit")
}

func displayVersion(){
	fmt.Println("yes (GNU coreutils) 1.0")
}

func main(){
	if err := rootCmd.Execute() ; err != nil {
		os.Exit(1)
	}
}

func printOutString(str string){

}