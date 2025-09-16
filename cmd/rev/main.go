package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	file := flag.String("F", "null", "File")
	flag.Parse()

	if *file != "null" {
		//reading from file
		fileInput, err := os.Open(*file)
		if err != nil {
			fmt.Println(err)
		}
		scanner := bufio.NewScanner(fileInput)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	} else {
		//Reading from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(reverseString(scanner.Text()))
		}

	}

}

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
