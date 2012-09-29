package main

import (
	"flag"
	"fmt"
	rl "github.com/igoralmeida/readline-go"
	"gonav"
	"runtime"
	"strings"
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

	runtime.GC()

	for {
		res := rl.ReadLine(&prompt)
		if res == nil {
			break
		}
		line := strings.TrimSpace(*res)
		if line != "" {
			rl.AddHistory(line)
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
