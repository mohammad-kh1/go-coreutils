package main


import (
	"os"
	"time"
	"fmt"

	"github.com/mohammad-kh1/go-coreutils/internal/utils"
	"github.com/spf13/cobra"

)


var (
	version bool
)

const HELP=`Usage: sleep NUMBER[SUFFIX]...
  or:  sleep OPTION
Pause for NUMBER seconds, where NUMBER is an integer or floating-point.
SUFFIX may be 's','m','h', or 'd', for seconds, minutes, hours, days.
With multiple arguments, pause for the sum of their values.

      --help        display this help and exit
      --version     output version information and exit

Report bugs to: bug-coreutils@gnu.org
GNU coreutils home page: <https://www.gnu.org/software/coreutils/>
General help using GNU software: <https://www.gnu.org/gethelp/>
Full documentation <https://www.gnu.org/software/coreutils/sleep>
or available locally via: info '(coreutils) sleep invocation'`

var rootCmd = &cobra.Command{
	Use:"Usage: sleep NUMBER[SUFFIX]...",
	Short:"Pause for NUMBER seconds, where NUMBER is an integer or floating-point.",
	Long: HELP,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command , args []string){
		if version{
			fmt.Println(utils.SLEEP_VERSION)
			return
		}

		for _ , v := range args{
			sleep(v)
		}

	},
}

func sleep(num string){
	i := 0
	rt := time.Second
	for ; i < len(num) ; i++{
		if num[i] >= '0' && num[i] <= '9' || num[i] == '.' {
			continue
		}
		i++
		break
	}
	i--
	if len(num)-i > 1 {
		fmt.Println("Invalid interval")
		os.Exit(1)
	}
	t := 0.0
	fmt.Sscanf(num , "%f" , &t)
	if len(num) - 1 == 1 {
		switch num[len(num)-1]{
			case 's' , 'S':
				rt = time.Second
			case 'm' , 'M':
				rt = time.Minute
			case 'h' , 'H':
				rt = time.Hour
			case 'd' , 'D':
				rt = time.Hour * 24

		}
	}

	rt = time.Duration(float64(rt) * t)
	time.Sleep(rt)
}

func init(){
	rootCmd.Flags().BoolVarP(&version , "version"  ,"v" , false , "output version information and exit")
}


func main(){
	if err := rootCmd.Execute() ; err != nil {
		os.Exit(1)
	}
}
