package main

import "strings"

type function struct {
	name  string
	usage string
	tags  []*tag
}

func newFunction(name string, comments []string) *function {
	tags := make([]*tag, 0)
	u := ""
	for _, comment := range comments {
		if strings.Contains(comment, name) {
			u = comment
		}
		if t := extractToTag(comment); t != nil {
			tags = append(tags, t)
		}
	}
	if len(tags) < 2 {
		return nil
	}
	return &function{
		name:  name,
		tags:  tags,
		usage: u,
	}
}
