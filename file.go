package main

import (
	"path"
	"strings"
)

type parseFile struct {
	// 包名
	pkg string
	// 路径名
	dir string
	// 文件名
	file string
	// 所有的方法
	funcs []*function
}

func newParseFile(filepath string) *parseFile {
	pkg := pkgPath(filepath)
	dir, filename := path.Split(filepath)
	return &parseFile{
		pkg:   pkg,
		dir:   dir,
		file:  filename,
		funcs: make([]*function, 0),
	}
}

func (p *parseFile) parse() {
	// 确认所有路由组
	for _, function := range p.funcs {
		r := &ginRouter{usage: function.usage, middlewares: make([]string, 0)}
		_, r.pkg = path.Split(p.pkg)
		g := defaultGroup
		for _, t := range function.tags {
			// 确认该function的tag有group
			if t.typ == group {
				if groupRouter[t.value] == nil {
					groupRouter[t.value] = make([]*ginRouter, 0)
				}
				g = t.value
			}
			// 确认路由路径
			if t.typ == router {
				r.path = t.value
			}
			// 确认方法
			if t.typ == method {
				r.method = t.value
			}
			// 确认中间件
			if t.typ == use {
				r.middlewares = strings.Split(t.value, ",")
				for i := range r.middlewares {
					r.middlewares[i] = r.pkg + "." + r.middlewares[i]
				}
			}
		}
		r.function = function
		groupRouter[g] = append(groupRouter[g], r)
	}
}
