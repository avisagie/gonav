package main

import (
	"flag"
	"fmt"
	"gonav"
	"strings"
	"readline"
)

var (
	prompt      = "\n> "
	root        = flag.String("root", ".", "The root directory to find go files to index")
	includeBody = flag.Bool("body", false, "Include the function body in the results")
)

func main() {
	fmt.Println("Starting up...")

	flag.Parse()
	types, functions := gonav.ProcessDir(*root, *includeBody)

	for {
		res, _ := readline.ReadLine(&prompt)
		if res == nil {
			break
		}
		line := strings.TrimSpace(*res)

		switch line {
		case "quit", "q", "exit":
			return
		}

		if line != "" {
			readline.AddHistory(line)
		}

		found := false

		byType, ok := types[line]
		if ok {
			found = true
			fmt.Println("Types:")
			fmt.Println("======")
			for _, f := range byType {
				fmt.Println(*f)
				fmt.Println("===========")
			}
		}

		byFunction, ok := functions[line]
		if ok {
			found = true
			fmt.Println("Functions:")
			fmt.Println("==========")
			for _, f := range byFunction {
				fmt.Println(*f)
				fmt.Println("===========")
			}
		}

		if !found {
			fmt.Println("Found nothing")
		}
	}
	fmt.Println()
}
