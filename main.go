package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "testdata/a.go", nil, 0)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		ident, ok := n.(*ast.Ident)
		if !ok {
			return true
		}

		if ident.Name != "Gopher" {
			return true
		}
		if ident.Obj.Kind != ast.Fun {
			return true
		}
		fdecl, ok := ident.Obj.Decl.(*ast.FuncDecl)
		if !ok {
			return true
		}
		if fdecl.Recv != nil {
			return true
		}
		fmt.Println(ident.Name)
		fmt.Println("discover!!!!")
		return true
	})

}
