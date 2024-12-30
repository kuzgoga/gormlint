package nullSafetyCheck

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"gormlint/common"
	"strings"
)

// NullSafetyAnalyzer todo: add URL for null safety analyzer rules
var NullSafetyAnalyzer = &analysis.Analyzer{
	Name: "GormNullSafety",
	Doc:  "reports problems with nullable fields with unsatisfied tag",
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

			if err := common.CheckUnnamedModel(*typeSpec); err != nil {
				pass.Reportf(structure.Pos(), err.Error())
				return false
			}

			for _, field := range structure.Fields.List {
				var structFieldName string
				if len(field.Names) == 0 {
					fieldType := common.ResolveBaseType(field.Type)
					if fieldType == nil {
						pass.Reportf(field.Pos(), "Failed to resolve model \"%s\" field type: %s", typeSpec.Name.Name, field.Type)
					} else {
						structFieldName = *fieldType
					}
				} else {
					structFieldName = field.Names[0].Name
				}
				if field.Tag != nil {
					tagWithoutQuotes := field.Tag.Value[1 : len(field.Tag.Value)-1]
					tagWithoutSemicolons := strings.ReplaceAll(tagWithoutQuotes, ";", ",")
					err := common.CheckFieldNullConsistency(*field, structFieldName, typeSpec.Name.Name, tagWithoutSemicolons)
					if err != nil {
						pass.Reportf(field.Pos(), err.Error())
						return false
					}
				}
			}
			return false
		})
	}
	return nil, nil
}
