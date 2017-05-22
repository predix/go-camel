package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/predix/go-camel/components/filereader"
	"github.com/predix/go-camel/components/printer"
	"github.com/predix/go-camel/core"
	"github.com/predix/go-camel/core/engine"
)

func main() {
	var filepath string

	flag.StringVar(&filepath, "f", "", "file to read")
	flag.Parse()

	if len(filepath) == 0 {
		fmt.Fprint(os.Stderr, "-f is a required argument")
		flag.Usage()
		os.Exit(1)
	}

	f, err := filereader.New(filepath)
	if err != nil {
		log.Fatalf("Failed to create a filereader for the file %q: %s", filepath, err)
	}

	con := printer.NewPrinter(os.Stdout)
	br := engine.NewRouteBuilder().From(f).To(con)
	c := core.NewContext()

	err = c.AddRoute(br)
	if err != nil {
		log.Fatalf("Failed to Add the route %q", err)
	}
	err = c.Start()
	if err != nil {
		log.Fatalf("Failed to Add the route %q", err)
	}
}
