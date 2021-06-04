package main

import "strings"

type function struct {
	name         string
	usage        string
	isMiddleware bool
	tags         []*tag
}

func newFunction(name string, comments []string) *function {
	tags := make([]*tag, 0)
	isMiddleware := false
	u := ""
	for _, comment := range comments {
		if strings.Contains(comment, name) {
			u = comment
		}
		if t := extractToTag(comment); t != nil {
			if t.typ == middleware {
				t.value = name
				isMiddleware = true
				tags = []*tag{t}
				break
			}
			tags = append(tags, t)
		}
	}
	return &function{
		name:         name,
		tags:         tags,
		usage:        u,
		isMiddleware: isMiddleware,
	}
}
