package main

import (
	"strings"
)

type ginRouter struct {
	middlewares []string
	function    *function
	path        string
	method      string
	pkg         string
	usage       string
}

func (r *ginRouter) middleware() string {
	if len(r.middlewares) == 0 {
		return ""
	}
	return strings.Join(r.middlewares, ",")
}

const (
	defaultGroup = "_"
)

var (
	groupRouter map[string][]*ginRouter
)

func init() {
	groupRouter = make(map[string][]*ginRouter)
	groupRouter[defaultGroup] = make([]*ginRouter, 0)
}

func (g *Generator) buildRouter(f []*parseFile) {
	for _, file := range f {
		file.parse()
	}
	for group, routers := range groupRouter {
		if group != defaultGroup {
			groupWrapper := findGroupByName(group)
			if groupWrapper == nil {
				continue
			}
			g.Printf("\n")
			g.Printf("%sGroup := engine.Group(%s)\n", firstLower(groupWrapper.name), groupWrapper.withPkgName())
			if len(groupWrapper.middlewares) > 0 {
				g.Printf("%sGroup.Use(%s)\n", firstLower(groupWrapper.name), groupWrapper.middleware())
			}
			g.Printf("{\n")
			for _, r := range routers {
				g.Printf("%s\n", r.usage)
				if len(r.middlewares) != 0 {
					g.Printf("%sGroup.%s(\"%s\", %s, %s.%s)\n", firstLower(groupWrapper.name), r.method, r.path, r.middleware(), r.pkg, r.function.name)
				} else {
					g.Printf("%sGroup.%s(\"%s\",%s.%s)\n", firstLower(groupWrapper.name), r.method, r.path, r.pkg, r.function.name)
				}
			}
			g.Printf("}\n")
			g.Printf("\n")
		} else {
			for _, r := range routers {
				g.Printf("%s\n", r.usage)
				if len(r.middlewares) != 0 {
					g.Printf("%s\n", r.usage)
					g.Printf("engine.%s(\"%s\", %s, %s.%s)\n", r.method, r.path, r.middleware(), r.pkg, r.function.name)
				} else {
					g.Printf("engine.%s(\"%s\",%s.%s)\n", r.method, r.path, r.pkg, r.function.name)
				}
			}
		}
	}
}
