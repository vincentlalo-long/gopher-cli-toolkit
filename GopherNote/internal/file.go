package internal

import (
	"fmt"
	"os"
	"strings"
)

func SaveNote(note *Note, data string) {
	f, err := os.OpenFile(note.FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Cannot open file")
	}
	defer f.Close()
	_, err = f.WriteString(data + "\n")
	if err != nil {
		fmt.Println("error cannot write file")
		return
	}
	info, _ := os.Stat(note.FileName)
	note.Size = uint64(info.Size())
	note.ModTime = info.ModTime()

}

func ListNotes(fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("No note found")
		return
	}
	content := string(data)
	lines := strings.Split(content, "\n")

	fmt.Println("------List note------")
	count := 0
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			count++
			fmt.Printf("%d %s\n", count, line)
		}
	}
	if count == 0 {
		fmt.Println("Empty Note")
	}
}
