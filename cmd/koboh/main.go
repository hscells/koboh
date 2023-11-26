package main

import (
	"fmt"
	"log"

	"github.com/hscells/koboh"
)

func main() {
	db, err := koboh.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	highlights, err := koboh.ExtractHighlights(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, highlight := range highlights {
		fmt.Printf("%s\n%s\n\n", highlight.Title, highlight.Highlight)
	}
}
