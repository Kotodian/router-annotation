package main

import (
	"strings"
	"testing"
)

func TestTrimSpace(t *testing.T) {
	comment := "                // Hello World           "
	comment = strings.Trim(comment, " ")
	//comment = strings.Replace(comment, " ", "", -1)
	t.Log(comment)
}
