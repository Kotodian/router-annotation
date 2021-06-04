package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	dir    = flag.String("dir", "/Users/linqiankai/go/src/router-annotation/examples", "input dir")
	output = flag.String("output", "/Users/linqiankai/go/src/router-annotation/api", "output dir")
)

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, "Usage of router-gen:\n")
	_, _ = fmt.Fprintf(os.Stderr, "	@router: router path\n")
	_, _ = fmt.Fprintf(os.Stderr, "	@method: http method\n")
	_, _ = fmt.Fprintf(os.Stderr, "	@group: router group\n")
	_, _ = fmt.Fprintf(os.Stderr, "	only support github.com/gin-gonic/gin\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()
	_, err := os.Stat(*dir)
	if err != nil {
		log.Println(err)
		flag.Usage()
		os.Exit(2)
	}
	_, err = os.Stat(*output)
	if err != nil {
		log.Println(err)
		flag.Usage()
		os.Exit(2)
	}
	g := newGenerator(*output)
	g.generate(*dir)
}
