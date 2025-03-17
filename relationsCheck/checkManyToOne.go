package relationsCheck

import (
	"go/types"
	"golang.org/x/tools/go/analysis"
	"gormlint/common"
	"strings"
)

func checkManyToOne(pass *analysis.Pass, nestedField common.Field, model common.Model, relatedModel common.Model) bool {
	foreignKey := nestedField.Tags.GetParamOr("foreignKey", model.Name+"Id")
	references := nestedField.Tags.GetParamOr("references", "Id")
	if !relatedModel.HasField(foreignKey) {
		pass.Reportf(
			nestedField.Pos,
			"Expected foreignKey `%s` in model `%s` for 1:M relation with model `%s`",
			foreignKey,
			relatedModel.Name,
			model.Name,
		)
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
		return true
	}
	foreignKeyType := types.ExprString(relatedModel.Fields[foreignKey].Type)
	referencesType := types.ExprString(model.Fields[references].Type)
	if !strings.Contains(foreignKeyType, "int") {
		// TODO: process UUID as foreign key type
		pass.Reportf(relatedModel.Fields[foreignKey].Pos, "Foreign key `%s` has invalid type", foreignKeyType)
		return true
	}

	if !strings.Contains(referencesType, "int") {
		pass.Reportf(model.Fields[references].Pos, "References key `%s` has invalid type", referencesType)
		return true
	}
	
	return false
}
