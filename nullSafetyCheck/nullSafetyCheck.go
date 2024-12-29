package nullSafetyCheck

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"gormlint/common"
	"strings"
)

var NullSafetyAnalyzer = &analysis.Analyzer{
	Name: "nullSafety",
	Doc:  "reports inconsistency of nullable values",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			typeSpec, ok := node.(*ast.TypeSpec)
			if !ok {
				return true
			}
			structure, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				return true
			}

			common.CheckUnnamedModel(*pass, *typeSpec)

			for _, field := range structure.Fields.List {
				common.CheckUnnamedField(*pass, typeSpec.Name.Name, *field)
				if field.Tag != nil {
					tagWithoutQuotes := field.Tag.Value[1 : len(field.Tag.Value)-1]
					tagWithoutSemicolons := strings.ReplaceAll(tagWithoutQuotes, ";", ",")
					common.CheckFieldNullConsistency(*pass, *field, typeSpec.Name.Name, tagWithoutSemicolons)
				} else {
					// TODO: check necessary tags for some fields
				}
			}
			return false
		})
	}
	return nil, nil
}
