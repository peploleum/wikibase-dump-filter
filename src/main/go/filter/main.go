package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	argsWithProg := os.Args
	log.Println("Executing filter")
	log.Println(argsWithProg)
	log.Println(strings.Repeat("▔", 65))
	log.Println(strings.Repeat("▔", 65))
	log.Println(strings.Repeat("▔", 65))
	log.Println("reading from stdin:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
