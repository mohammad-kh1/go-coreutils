package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	file := flag.String("file", "null", "file")
	del := flag.String("d", "", "delimiter")
	field := flag.Int("f", 0, "field")
	flag.Parse()

	if *file != "null" {
		//reading from file
		fileInput, err := os.Open(*file)
		defer fileInput.Close()
		if err != nil {
			fmt.Println(err)
		}
		scanner := bufio.NewScanner(fileInput)
		for scanner.Scan() {
			line := strings.Split(scanner.Text(), *del)
			if *field >= 0 && *field < len(os.Args) {
				fmt.Println(line[*field])
			}
		}
	} else {
		//reading from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := strings.Split(scanner.Text(), *del)
			if *field >= 0 && *field < len(os.Args) {
				fmt.Println(line[*field])
			}
		}
	}
}
