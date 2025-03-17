package relationsCheck

import (
	"go/token"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"gormlint/common"
	"strings"
)

var alreadyReported = make(map[token.Pos]bool)

func checkManyToOne(pass *analysis.Pass, nestedField common.Field, model common.Model, relatedModel common.Model) bool {
	/* Return true, if found problems */

	foreignKey := nestedField.Tags.GetParamOr("foreignKey", model.Name+"Id")
	references := nestedField.Tags.GetParamOr("references", "Id")

	if alreadyReported[nestedField.Pos] {
		return true
	}

	if !relatedModel.HasField(foreignKey) {
		pass.Reportf(
			nestedField.Pos,
			"Expected foreignKey `%s` in model `%s` for 1:M relation with model `%s`",
			foreignKey,
			relatedModel.Name,
			model.Name,
		)
		alreadyReported[nestedField.Pos] = true
		return true
	}

	if !model.HasField(references) {
		pass.Reportf(
			nestedField.Pos,
			"Expected references `%s` in model `%s` for 1:M relation with model `%s`",
			references,
			model.Name,
			relatedModel.Name,
		)
		alreadyReported[nestedField.Pos] = true
		return true
	}

	foreignKeyField := relatedModel.Fields[foreignKey]
	referencesField := model.Fields[references]
	foreignKeyType := types.ExprString(foreignKeyField.Type)
	referencesType := types.ExprString(referencesField.Type)

	if alreadyReported[foreignKeyField.Pos] || alreadyReported[referencesField.Pos] {
		return true
	}

	if !strings.Contains(foreignKeyType, "int") && !alreadyReported[foreignKeyField.Pos] {
		// TODO: process UUID as foreign key type
		pass.Reportf(foreignKeyField.Pos, "Foreign key `%s` has invalid type", foreignKeyType)
		alreadyReported[foreignKeyField.Pos] = true
		return true
	}

	if !strings.Contains(referencesType, "int") && !alreadyReported[referencesField.Pos] {
		pass.Reportf(referencesField.Pos, "References key `%s` has invalid type", referencesType)
		alreadyReported[referencesField.Pos] = true
		return true
	}

	return false
}
