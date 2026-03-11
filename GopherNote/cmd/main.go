package main

import (
	"GopherNote/internal"
	"flag"
	"fmt"
)

func main() {
	addPtr := flag.String("add", "", "note text")
	filePtr := flag.String("file", "noteAll.txt", "name file note")
	listPtr := flag.Bool("list", false, "list of all list with id")
	flag.Parse()
	if *listPtr {
		internal.ListNotes(*filePtr)
		return
	}
	if *addPtr != "" {
		myNote := internal.Note{FileName: *filePtr}
		internal.SaveNote(&myNote, *addPtr)
		fmt.Printf("Success! File size %d , Mod Time : %s\n", myNote.Size, myNote.ModTime)
	} else {
		fmt.Println("Use -add to add text to note")
	}

}
