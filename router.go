package main

import "path"

type ginRouter struct {
	middleware string
	function   *function
	path       string
	method     string
	pkg        string
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
			_, groupName := path.Split(group)
			g.Printf("%sGroup := engine.Group(\"%s\")\n", groupName, group)
			g.Printf("{\n")
			for _, r := range routers {
				if r.middleware != "" {
					//todo: 加入中间件
				} else {
					g.Printf("%sGroup.%s(\"%s\",%s.%s)\n", groupName, r.method, r.path, r.pkg, r.function.name)
				}
			}
			g.Printf("}\n")
		} else {
			for _, r := range routers {
				g.Printf("engine.%s(\"%s\",%s.%s)\n", r.method, r.path, r.pkg, r.function.name)
			}
		}
	}
}
