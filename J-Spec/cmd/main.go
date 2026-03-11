package main

import (
	"flag"
	"fmt"
	"os"

	"J-Spec/internal/formatter"
	"J-Spec/internal/parser"
	"J-Spec/internal/processor"
)

func main() {
	// 1.  Flags
	filePtr := flag.String("file", "", "path file JSON")
	pathPtr := flag.String("path", "", "Path key (ex: user.name)")
	filterPtr := flag.String("filter", "", "Condition filter (example: 'tasks > 5')")
	prettyPtr := flag.Bool("pretty", true, "color output")
	flag.Parse()

	var input *os.File
	var err error

	if *filePtr != "" {
		input, err = parser.OpenFile(*filePtr)
		if err != nil {
			fmt.Printf(" Open file error : %v\n", err)
			return
		}
		defer input.Close()
	} else {
		// ex cat data.json | jspec -path="name"
		input = os.Stdin
	}

	// 3. Parse JSON
	data, err := parser.LoadJson(input)
	if err != nil {
		fmt.Printf("Error parse: %v\n", err)
		return
	}

	// 4. Path Selector
	result := data
	if *pathPtr != "" {
		result, err = processor.GetValueByPath(data, *pathPtr)
		if err != nil {
			fmt.Printf(" Error Selector: %v\n", err)
			return
		}
	}

	// 5. Filter
	if *filterPtr != "" {
		result, err = processor.FilterData(result, *filterPtr)
		if err != nil {
			fmt.Printf("Error Filter: %v\n", err)
			return
		}
	}

	// 6. final result
	formatter.Print(result, *prettyPtr)
}
