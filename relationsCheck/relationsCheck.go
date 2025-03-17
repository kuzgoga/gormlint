package relationsCheck

import (
	"fmt"
	"golang.org/x/tools/go/analysis"
	"gormlint/common"
	"slices"
)

// RelationsAnalyzer todo: add URL for rule
var RelationsAnalyzer = &analysis.Analyzer{
	Name: "GormRelationsCheck",
	Doc:  "report about invalid relations in models",
	Run:  run,
}

func CheckTypesOfM2M(pass *analysis.Pass, modelName string, relatedModelName string, relationName string, reference common.Field, backReference common.Field) bool {
	if !common.IsSlice(reference.Type) {
		pass.Reportf(reference.Pos, "M2M relation `%s` with bad type `%s` (should be a slice)", relationName, reference.Type)
		return true
	}
	if !common.IsSlice(backReference.Type) {
		pass.Reportf(backReference.Pos, "M2M relation `%s` with bad type `%s` (should be a slice)", relationName, backReference.Type)
		return true
	}

	referenceBaseType := common.ResolveBaseType(reference.Type)
	if referenceBaseType == nil {
		pass.Reportf(reference.Pos, "Failed to resolve field type: `%s`", reference.Type)
		return true
	}
	backReferenceBaseType := common.ResolveBaseType(backReference.Type)
	if backReferenceBaseType == nil {
		pass.Reportf(reference.Pos, "Failed to resolve type: `%s`", reference.Type)
		return true
	}

	if *backReferenceBaseType != modelName {
		pass.Reportf(backReference.Pos, "Invalid type `%s` in M2M relation (use []*%s or self-reference)", *backReferenceBaseType, modelName)
		return true
	}

	if *referenceBaseType != relatedModelName {
		pass.Reportf(reference.Pos, "Invalid type `%s` in M2M relation (use []*%s or self-reference)", *referenceBaseType, relatedModelName)
		return true
	}
	return false
}

func CheckMany2Many(pass *analysis.Pass, models map[string]common.Model) {
	// TODO: unexpected duplicated relations
	var knownModels []string
	for _, model := range models {
		for _, field := range model.Fields {
			m2mRelation := field.Tags.GetParam("many2many")
			if m2mRelation != nil {
				relatedModel := common.GetModelFromType(field.Type, models)
				if relatedModel == nil {
					pass.Reportf(field.Pos, "Failed to resolve related model type")
					continue
				}

				backReference := common.FindBackReferenceInM2M(m2mRelation.Value, *relatedModel)
				if backReference != nil {
					if slices.Contains(knownModels, relatedModel.Name) {
						continue
					} else {
						knownModels = append(knownModels, model.Name)
						knownModels = append(knownModels, relatedModel.Name)
					}
					if CheckTypesOfM2M(pass, model.Name, relatedModel.Name, m2mRelation.Value, field, *backReference) {
						continue
					}
					// TODO: check foreign key and references
					fmt.Printf("Found M2M relation between \"%s\" and \"%s\"\n", model.Name, relatedModel.Name)
					if CheckCascadeDelete(pass, field) {
						continue
					}
				} else {
					// Check self-reference
					if model.Name == relatedModel.Name {
						CheckTypesOfM2M(pass, model.Name, relatedModel.Name, m2mRelation.Value, field, field)
					} else {
						if !relatedModel.HasPrimaryKey() {
							fmt.Printf("%#v\n", relatedModel)
							pass.Reportf(field.Pos, "Can't build M2M relation `%s`, primary key on `%s` model is absont", m2mRelation.Value, relatedModel.Name)
							continue
						}
					}
					// Here you can forbid M2M relations without back-reference
					// TODO: process m2m without backref
					if CheckCascadeDelete(pass, field) {
						continue
					}
				}
			} else {
				if common.IsSlice(field.Type) {
					relatedModel := common.GetModelFromType(field.Type, models)
					if relatedModel == nil {
						pass.Reportf(field.Pos, "Failed to resolve related model type")
						continue
					}
					if checkManyToOne(pass, field, model, *relatedModel) {
						continue
					}
					if CheckCascadeDelete(pass, field) {
						continue
					}
				}
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
