package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/vibhavp/dasm-go/disasm"
	"github.com/vibhavp/dasm-go/read"
)

var file = flag.String("file", "", "File to run")

func init() {
	flag.Parse()
}

func main() {
	if *file == "" {
		flag.Usage()
		os.Exit(1)
	}
	b, err := read.FromFile(*file)
	if err != nil {
		log.Fatal(err)
	}

	out, err := disasm.Dump(b)
	fmt.Printf("Code: %s\n", out)
}
