package relationsCheck

import (
	"fmt"
	"golang.org/x/tools/go/analysis"
	"gormlint/common"
	"slices"
)

// RelationsAnalyzer todo: add URL for rule
var RelationsAnalyzer = &analysis.Analyzer{
	Name: "GormReferencesCheck",
	Doc:  "report about invalid references in models",
	Run:  run,
}

func CheckTypesOfM2M(pass *analysis.Pass, modelName string, relatedModelName string, relationName string, reference common.Field, backReference common.Field) {
	if !common.IsSlice(reference.Type) {
		pass.Reportf(reference.Pos, "M2M relation `%s` with bad type `%s` (should be a slice)", relationName, reference.Type)
		return
	}
	if !common.IsSlice(backReference.Type) {
		pass.Reportf(backReference.Pos, "M2M relation `%s` with bad type `%s` (should be a slice)", relationName, backReference.Type)
		return
	}

	referenceBaseType := common.ResolveBaseType(reference.Type)
	if referenceBaseType == nil {
		pass.Reportf(reference.Pos, "Failed to resolve field type: `%s`", reference.Type)
		return
	}
	backReferenceBaseType := common.ResolveBaseType(backReference.Type)
	if backReferenceBaseType == nil {
		pass.Reportf(reference.Pos, "Failed to resolve type: `%s`", reference.Type)
		return
	}

	if *backReferenceBaseType != modelName {
		pass.Reportf(backReference.Pos, "Invalid type `%s` in M2M relation (use []*%s or self-reference)", *backReferenceBaseType, modelName)
		return
	}

	if *referenceBaseType != relatedModelName {
		pass.Reportf(reference.Pos, "Invalid type `%s` in M2M relation (use []*%s or self-reference)", *referenceBaseType, relatedModelName)
	}
}

func CheckMany2Many(pass *analysis.Pass, models map[string]common.Model) {
	// TODO: unexpected duplicated relations
	var knownModels []string
	for _, model := range models {
		for _, field := range model.Fields {
			m2mRelation := field.GetParam("many2many")
			if m2mRelation != nil {
				relatedModel := common.GetModelFromType(field.Type, models)
				if relatedModel == nil {
					pass.Reportf(field.Pos, "Failed to resolve related model type")
					return
				}

				backReference := common.FindBackReferenceInM2M(m2mRelation.Value, *relatedModel)
				if backReference != nil {
					if slices.Contains(knownModels, relatedModel.Name) {
						continue
					} else {
						knownModels = append(knownModels, model.Name)
						knownModels = append(knownModels, relatedModel.Name)
					}
					CheckTypesOfM2M(pass, model.Name, relatedModel.Name, m2mRelation.Value, field, *backReference)
					// TODO: check foreign key and references
					fmt.Printf("Found M2M relation between \"%s\" and \"%s\"\n", model.Name, relatedModel.Name)
				} else {
					// Here you can forbid M2M relations without back-reference
					// TODO: process m2m without backref
				}
			} else {
				// TODO: check [] and process m:1
			}
		}
	}
}

func run(pass *analysis.Pass) (any, error) {
	models := make(map[string]common.Model)
	common.ParseModels(pass, &models)
	CheckMany2Many(pass, models)

	return nil, nil
}
