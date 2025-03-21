package relationsCheck

import (
	"fmt"
	"github.com/kuzgoga/gormlint/common"
	"golang.org/x/tools/go/analysis"
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
	var processedRelations []string

	for _, model := range models {
		for _, field := range model.Fields {
			m2mRelation := field.Tags.GetParam("many2many")
			if m2mRelation != nil {
				if slices.Contains(processedRelations, m2mRelation.Value) {
					continue
				}
				processedRelations = append(processedRelations, m2mRelation.Value)

				relatedModel := common.GetModelFromType(field.Type, models)
				if relatedModel == nil {
					pass.Reportf(field.Pos, "Failed to resolve related model type")
					continue
				}

				backReference := common.FindBackReferenceInM2M(m2mRelation.Value, *relatedModel)
				if backReference != nil {
					if CheckTypesOfM2M(pass, model.Name, relatedModel.Name, m2mRelation.Value, field, *backReference) {
						continue
					}
					// Проверка каскадного удаления и других параметров
					if CheckCascadeDelete(pass, field) {
						continue
					}
				} else {
					// Обработка самоссылки
					if model.Name == relatedModel.Name {
						if CheckTypesOfM2M(pass, model.Name, relatedModel.Name, m2mRelation.Value, field, field) {
							continue
						}
					} else {
						pass.Reportf(field.Pos, "M2M relation `%s` missing back-reference in model `%s`", m2mRelation.Value, relatedModel.Name)
					}
					if CheckCascadeDelete(pass, field) {
						continue
					}
				}
			}
		}
	}
}

func CheckOneToMany(pass *analysis.Pass, models map[string]common.Model) {
	for _, model := range models {
		for _, field := range model.Fields {
			if common.IsSlice(field.Type) {
				continue
			}
			if field.Tags.HasParam("many2many") {
				continue
			}

			baseType := common.ResolveBaseType(field.Type)
			if baseType == nil {
				pass.Reportf(field.Pos, "Failed to resolve field base type: `%s`", field.Type)
				continue
			}
			relatedModel := common.GetModelFromType(field.Type, models)
			if relatedModel == nil {
				continue
			}

			foundOneToMany := isOneToMany(pass, model, *relatedModel)
			if foundOneToMany {
				fmt.Printf("Found 1:M relation in model `%s` with model `%s`\n", model.Name, *baseType)
			}

			foundBelongsTo := IsBelongsTo(field, model, *relatedModel)
			hasOne := IsHasOne(field, model, *relatedModel)

			if !foundOneToMany {
				if foundBelongsTo {
					fmt.Printf("`%s` belongs `%s`\n", *baseType, model.Name)
					if CheckCascadeDelete(pass, field) {
						return
					}
				} else if hasOne {
					fmt.Printf("`%s` has one `%s` \n", model.Name, relatedModel.Name)
					if CheckCascadeDelete(pass, field) {
						return
					}
				} else {
					pass.Reportf(field.Pos, "Invalid relation in field `%s`", field.Name)
				}
			}
		}
	}
}

func run(pass *analysis.Pass) (any, error) {
	models := make(map[string]common.Model)
	common.ParseModels(pass, &models)
	CheckMany2Many(pass, models)
	CheckOneToMany(pass, models)
	return nil, nil
}
