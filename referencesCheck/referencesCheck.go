package referencesCheck

import (
	"github.com/fatih/structtag"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"gormlint/common"
	"strings"
)

// ReferenceAnalyzer todo: add URL for rule
var ReferenceAnalyzer = &analysis.Analyzer{
	Name: "gormReferencesCheck",
	Doc:  "report about invalid references in models",
	Run:  run,
}

var models map[string]common.Model

func init() {
	models = make(map[string]common.Model)
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

			var model common.Model
			model.Name = typeSpec.Name.Name
			model.Comment = typeSpec.Comment.Text()
			model.Position = structure.Pos()
			model.Fields = make(map[string]common.Field)

			for _, field := range structure.Fields.List {
				var structField common.Field
				if err := common.CheckUnnamedField(typeSpec.Name.Name, *field); err != nil {
					pass.Reportf(field.Pos(), err.Error())
					return false
				}
				structField.Name = field.Names[0].Name
				structField.Position = field.Pos()
				structField.Comment = field.Comment.Text()
				structField.Type = field.Type
				if field.Tag != nil {
					structField.Tags = &field.Tag.Value

					tags, err := structtag.Parse(common.NormalizeStructTags(field.Tag.Value))
					if err != nil {
						pass.Reportf(field.Pos(), "Invalid structure tag: %s\n", err)
						return false
					}
					if tags != nil {
						gormTag, parseErr := tags.Get("gorm")
						if gormTag != nil && parseErr == nil {
							gormTag.Options = append([]string{gormTag.Name}, gormTag.Options...)
							for _, opt := range gormTag.Options {
								if strings.Contains(opt, ":") {
									structField.Params = append(structField.Options, opt)
								} else {
									structField.Options = append(structField.Options, opt)
								}
							}
						}
						if parseErr != nil {
							pass.Reportf(field.Pos(), "Invalid structure tag: %s\n", parseErr)
							return false
						}
					}

					model.Fields[structField.Name] = structField
				}

				models[model.Name] = model
			}
			return false
		})
	}
	return nil, nil
}
