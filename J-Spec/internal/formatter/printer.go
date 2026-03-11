package formatter

import (
	"encoding/json"
	"fmt"
)

const (
	ColorCyan   = "\033[36m"
	ColorReset  = "\033[0m"
	ColorYellow = "\033[33m"
)

func Print(data any, pretty bool) {
	var b []byte
	var err error

	if pretty {
		b, err = json.MarshalIndent(data, "", "  ")
	} else {
		b, err = json.Marshal(data)
	}

	if err != nil {
		fmt.Printf("Error format: %v\n", err)
		return
	}

	fmt.Printf("%s---J-Spec Result ---%s\n", ColorCyan, ColorReset)
	fmt.Println(string(b))
	fmt.Printf("%s------------------------%s\n", ColorCyan, ColorReset)
}
