package common

import (
	"github.com/fatih/structtag"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"strings"
)

func ParseModels(pass *analysis.Pass, models *map[string]Model) {
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

			if err := CheckUnnamedModel(*typeSpec); err != nil {
				pass.Reportf(structure.Pos(), err.Error())
				return false
			}

			var model Model
			model.Name = typeSpec.Name.Name
			model.Comment = typeSpec.Comment.Text()
			model.Pos = structure.Pos()
			model.Fields = make(map[string]Field)

			for _, field := range structure.Fields.List {
				var structField Field

				if len(field.Names) == 0 {
					fieldType := ResolveBaseType(field.Type)
					if fieldType == nil {
						pass.Reportf(field.Pos(), "Failed to resolve model \"%s\" field type: %s", model.Name, field.Type)
					} else {
						structField.Name = *fieldType
					}
				} else {
					structField.Name = field.Names[0].Name
				}

				structField.Pos = field.Pos()
				structField.Comment = field.Comment.Text()
				structField.Type = field.Type

				if field.Tag != nil {
					structField.Tags = &field.Tag.Value

					tags, err := structtag.Parse(NormalizeStructTags(field.Tag.Value))
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
									structField.Params = append(structField.Params, opt)
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
				}
				model.Fields[structField.Name] = structField
				(*models)[model.Name] = model
			}
			return false
		})
	}
}
