package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
)

func handwrittenMethods(dir string) (map[string]bool, error) {
	paths, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		return nil, err
	}
	methods := make(map[string]bool)
	for _, p := range paths {
		n := filepath.Base(p)
		if strings.HasPrefix(n, "gen_") || strings.HasSuffix(n, "_test.go") {
			continue
		}
		f, err := parser.ParseFile(token.NewFileSet(), p, nil, 0)
		if err != nil {
			return nil, err
		}
		for _, d := range f.Decls {
			fn, ok := d.(*ast.FuncDecl)
			if !ok || fn.Recv == nil {
				continue
			}
			t := fn.Recv.List[0].Type
			if p, ok := t.(*ast.StarExpr); ok {
				t = p.X
			}
			if id, ok := t.(*ast.Ident); ok {
				methods[id.Name+"."+fn.Name.Name] = true
			}
		}
	}
	return methods, nil
}
