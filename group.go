package main

import (
	"path"
	"strings"
)

var (
	groups []*ginGroup
)

func init() {
	groups = make([]*ginGroup, 0)
}

type ginGroup struct {
	pkg         string
	middlewares []string
	name        string
	tags        []*tag
	value       string
}

func (g *ginGroup) middleware() string {
	if len(g.middlewares) == 0 {
		return ""
	}
	return strings.Join(g.middlewares, ",")
}

func newGinGroup(name, value, pkgPath string, comments []string) *ginGroup {
	_, pkg := path.Split(pkgPath)
	middlewares := make([]string, 0)
	g := &ginGroup{
		pkg:   pkg,
		name:  name,
		value: value,
	}
	for _, comment := range comments {
		if t := extractToTag(comment); t != nil {
			if t.typ == use {
				middlewares = strings.Split(t.value, ",")
			}
		}
	}
	if len(middlewares) > 0 {
		g.middlewares = make([]string, len(middlewares))
		copy(g.middlewares, middlewares)
	}
	return g
}

func findGroupByName(name string) *ginGroup {
	for _, g := range groups {
		if g.name == name {
			return g
		}
	}
	return nil
}

func findGroupByValue(value string) *ginGroup {
	for _, g := range groups {
		if g.value == value {
			return g
		}
	}
	return nil
}

func (g *ginGroup) withPkgName() string {
	return g.pkg + "." + g.name
}
