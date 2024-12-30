package common

import (
	"go/ast"
)

func ResolveBaseType(expr ast.Expr) *string {
	switch e := expr.(type) {
	case *ast.Ident:
		return &e.Name
	case *ast.StarExpr:
		return ResolveBaseType(e.X)
	case *ast.ArrayType:
		return ResolveBaseType(e.Elt)
	case *ast.SelectorExpr:
		return ResolveBaseType(e.X)
	case *ast.ParenExpr:
		return ResolveBaseType(e.X)
	}
	return nil
}
