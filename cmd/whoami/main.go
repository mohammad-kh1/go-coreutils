package main

import (
	"os/user"
	"fmt"
	"os"

	"github.com/mohammad-kh1/go-coreutils/internal/utils"
	"github.com/spf13/cobra"

)

var (
	version bool
)

const HELP=`Usage: whoami [OPTION]...
Print the user name associated with the current effective user ID.
Same as id -un.

      --help        display this help and exit
      --version     output version information and exit

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/whoami>
or available locally via: info '(coreutils) whoami invocation'`

var rootCmd = &cobra.Command{
	Use:"Usage: whoami [OPTION]...",
	Short:"Print the user name associated with the current effective user ID.",
	Long:HELP,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command , args []string){
		if version{
			fmt.Println(utils.WHOAMI_VERSION)
			return
		}
		user , err := user.Current()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(user.Name)
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
