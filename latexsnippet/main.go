package main

import (
	"os"
	"log"
	"github.com/nounoursheureux/latexsnippet"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: " + os.Args[0] + " <LaTeX snippet> <filename>")
	}

	var err = latexsnippet.RenderSnippet(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
}
