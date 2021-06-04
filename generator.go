package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Generator struct {
	// 缓冲区
	buf bytes.Buffer
	// 所有文件
	files []string
	// 输出路径
	output string
}

func newGenerator(output string) *Generator {
	return &Generator{
		files:  make([]string, 0),
		output: output,
	}
}

func (g *Generator) generate(dir string) {
	g.parseDir(dir)
	files := g.parseFiles()
	if files == nil {
		return
	}
	pkgPath := make([]string, 0)
	for _, file := range files {
		pkgPath = append(pkgPath, file.pkg)
	}
	g.buildPrepare(pkgPath)
	g.buildRouter(files)
	g.buildEnd()
	src := g.format()
	// 确认生成文件
	outputFile := path.Join(g.output, "register.go")
	// 输出到指定文件
	err := os.WriteFile(outputFile, src, 0644)
	if err != nil {
		panic(err)
	}
	// 重置
	g.reset()
}
func (g *Generator) Printf(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(&g.buf, format, args...)
}

func (g *Generator) parseDir(dir string) {
	_ = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			if filepath.Ext(d.Name()) == ".go" {
				g.files = append(g.files, path)
			}
		}
		return nil
	})
}

func (g *Generator) parseFiles() []*parseFile {
	if len(g.files) == 0 {
		return nil
	}
	parseFiles := make([]*parseFile, 0)
	for _, file := range g.files {
		set := token.NewFileSet()
		// 生成解析文件
		parseFile := newParseFile(file)
		// 解析该文件的所有注释
		f, _ := parser.ParseFile(set, file, nil, parser.ParseComments)
		funcs := make([]*function, 0)
		ast.Inspect(f, func(node ast.Node) bool {
			switch t := node.(type) {
			case *ast.FuncDecl:
				if t.Doc == nil || t.Doc.List == nil {
					return true
				}
				_comments := t.Doc.List
				comments := make([]string, 0)
				for _, comment := range _comments {
					comments = append(comments, comment.Text)
				}
				function := newFunction(t.Name.Name, comments)
				funcs = append(funcs, function)
			}
			return true
		})
		if funcs == nil {
			return nil
		}
		parseFile.funcs = funcs
		parseFiles = append(parseFiles, parseFile)
	}
	return parseFiles
}

func (g *Generator) buildPrepare(pkgPath []string) {
	output := g.output[strings.LastIndex(g.output, "/")+1:]
	g.Printf("// Package %s Code generated by \"router-annotation\";DO NOT EDIT.\n", output)
	g.Printf("package %s\n", output)
	g.Printf("import (\n")
	g.Printf("\"github.com/gin-gonic/gin\"\n")
	for _, p := range pkgPath {
		g.Printf("\"%s\"\n", p)
	}
	g.Printf(")\n")
	g.Printf("var engine *gin.Engine\n")
	g.Printf("func init() {\n")
	g.Printf("engine = gin.Default()\n")
}

func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}

func (g *Generator) buildEnd() {
	g.Printf("}\n")
	g.Printf("func GinEngine() *gin.Engine {\n")
	g.Printf("return engine \n")
	g.Printf("}\n")
}

func (g *Generator) reset() {
	g.buf.Reset()
}