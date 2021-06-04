package main

import "flag"

var (
	dir    = flag.String("dir", "/Users/linqiankai/go/src/router-annotation/examples", "input dir")
	output = flag.String("output", "/Users/linqiankai/go/src/router-annotation/api", "output dir")
)

func main() {
	flag.Parse()
	g := newGenerator(*output)
	g.generate(*dir)
}
