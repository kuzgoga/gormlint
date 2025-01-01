package common

import "go/ast"

func IsSlice(expr ast.Expr) bool {
	arrayType, ok := expr.(*ast.ArrayType)
	if !ok {
		return false
	}
	return arrayType.Len == nil
}
