package referencesCheck

import (
	"golang.org/x/tools/go/analysis"
	"gormlint/common"
	"strings"
)

// ReferenceAnalyzer todo: add URL for rule
var ReferenceAnalyzer = &analysis.Analyzer{
	Name: "GormReferencesCheck",
	Doc:  "report about invalid references in models",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	models := make(map[string]common.Model)
	common.ParseModels(pass, &models)

	for _, model := range models {
		for _, field := range model.Fields {
			for _, param := range field.Params {
				pair := strings.Split(param, ":")
				paramName := pair[0]
				paramValue := pair[1]

				if paramName == "reference" {
					pass.Reportf(field.Pos, "Typo in tag: \"reference\" instead of verb \"references\"")
				}
				if paramName == "references" {
					fieldType := common.ResolveBaseType(field.Type)

					if fieldType == nil {
						pass.Reportf(field.Pos, "Failed to process references check. Cannot resolve type \"%s\" in field \"%s\"", field.Type, field.Name)
						return nil, nil
					}

					relatedModel, modelExists := models[*fieldType]

					if modelExists {
						_, fieldExists := relatedModel.Fields[paramValue]
						if !fieldExists {
							pass.Reportf(field.Pos, "Related field \"%s\" doesn't exist on model \"%s\"", paramValue, relatedModel.Name)
						}
					} else {
						pass.Reportf(field.Pos, "Related model \"%s\" doesn't exist", *fieldType)
					}
				}
			}
		}
	}
	return nil, nil
}
