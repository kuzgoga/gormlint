package common

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func CheckUnnamedModel(pass analysis.Pass, typeSpec ast.TypeSpec) {
	if typeSpec.Name == nil {
		pass.Reportf(typeSpec.Pos(), "Unnamed model\n")
	}
}

func CheckUnnamedField(pass analysis.Pass, structName string, field ast.Field) {
	if len(field.Names) == 0 {
		pass.Reportf(field.Pos(), "Struct \"%s\" has unnamed field", structName)
	}
}
