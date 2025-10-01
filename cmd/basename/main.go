package main

import (
	"strings"
	"fmt"
	"os"

	"github.com/mohammad-kh1/go-coreutils/internal/utils"

	"github.com/spf13/cobra"

)

var (
	multiple bool
	suffix string
	zero	bool
	version bool
)

const seperator = "/"

const NoArgsMessage = `basename: missing operand
Try 'basename --help' for more information.`

const HELP = `Usage: basename NAME [SUFFIX]
  or:  basename OPTION... NAME...
Print NAME with any leading directory components removed.
If specified, also remove a trailing SUFFIX.

Mandatory arguments to long options are mandatory for short options too.
  -a, --multiple       support multiple arguments and treat each as a NAME
  -s, --suffix=SUFFIX  remove a trailing SUFFIX; implies -a
  -z, --zero           end each output line with NUL, not newline
      --help        display this help and exit
      --version     output version information and exit

Examples:
  basename /usr/bin/sort          -> "sort"
  basename include/stdio.h .h     -> "stdio"
  basename -s .h include/stdio.h  -> "stdio"
  basename -a any/str1 any/str2   -> "str1" followed by "str2"

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/basename>
or available locally via: info '(coreutils) basename invocation'`


var rootCmd = &cobra.Command{

	Use:"Usage: basename NAME [SUFFIX] or:  basename OPTION... NAME...",
	Short:"Print NAME with any leading directory components removed.",
	Long: HELP,
	Args: cobra.ArbitraryArgs,
	Run : func(cmd *cobra.Command , args []string){

		if version{
			fmt.Println(utils.BASENAME_VERSION)
			return
		}

		if len(args) == 0 {
			fmt.Println(NoArgsMessage)
			return
		}
		output := splitPath(args)

		if suffix != "" {
			for i := 0 ; i < len(output) ; i++{
				output[i] = strings.Split(output[i] , suffix)[0]
			}
		}
		if zero == true{
			if multiple == false{
				fmt.Print(output[0])
				return
			}else{
				for _ , value := range output {
					fmt.Print(value)
				}
				return
			}
		}

		if multiple == false{
			fmt.Println(output[0])
			return
		}
		for _ , path := range output {
			fmt.Println(path)
		}

	},
}

func splitPath(basename []string) []string{
	var output []string 
	for _ , path := range basename {
		base := strings.Split(path , seperator)
		output = append(output , base[len(base)-1])
	}
	return output
}

func init(){
	rootCmd.Flags().BoolVarP(&multiple , "multiple" , "a" , false , "support multiple arguments and treat each as a NAME")
	rootCmd.Flags().StringVarP(&suffix , "suffix" , "s" , "" , "remove a trailing SUFFIX; implies -a")
	rootCmd.Flags().BoolVarP(&zero , "zero" , "z", false , "end each output line with NUL, not newline")
	rootCmd.Flags().BoolVarP(&version , "version" , "v", false , "output version information and exit")

}

func main(){
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}