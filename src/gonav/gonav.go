// Parse go source files to find functions. Indexes them in simple
// maps to allow quickly finding functions by a specific name, and the
// functions on a type. Does not work through embedding yet.
//
package gonav

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func Parse(expr string) {
	e, err := parser.ParseExpr(expr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	printer.Fprint(os.Stdout, token.NewFileSet(), e)
	fmt.Println()
}

func Walk(root string) (gofiles []string) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".go") && !info.IsDir() {
			gofiles = append(gofiles, path)
		}
		return nil
	})
	return gofiles
}

func ProcessDir(root string, includeBody bool) (types, functions map[string][]*string) {
	files := Walk(root)
	fset := token.NewFileSet()

	types = make(map[string][]*string)
	functions = make(map[string][]*string)

	for _, fn := range files {
		a, err := parser.ParseFile(fset, fn, nil, parser.ParseComments)
		if err != nil {
			fmt.Println("Could not parse", fn, "-", err)
			continue
		}

		processAst(a, fset, types, functions, includeBody)
	}

	return types, functions
}

var buf = &bytes.Buffer{}

// walk the ast and extract function declarations. Build up the
// indexes.
func processAst(a *ast.File, fset *token.FileSet, types, functions map[string][]*string, includeBody bool) {
	ast.Inspect(a, func(node ast.Node) bool {
		buf.Reset()
		switch x := node.(type) {
		case *ast.FuncDecl:
			pos := fset.File(node.Pos())
			fmt.Fprintf(buf, "%s:%d: \n", pos.Name(), pos.Line(node.Pos()))
			if !includeBody {
				x.Body = nil
			}
			printer.Fprint(buf, fset, x)

			function := buf.String()
			add(functions, x.Name.Name, &function)
			if x.Recv != nil {
				buf.Reset()
				printer.Fprint(buf, fset, x.Recv.List[0].Type)
				thetype := strings.Trim(buf.String(), "*")
				add(types, thetype, &function)
			}
		}
		return true
	})
}

func add(things map[string][]*string, key string, value *string) {
	things[key] = append(things[key], value)
}
