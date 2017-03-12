package main

import (
	"flag"
	"log"
	"os"

	"github.com/vibhavp/dasm-go/read"
	"github.com/vibhavp/dasm-go/vm"
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

	vm.Run(b.Bytecode, b.MaxStackDepth, true)
}
