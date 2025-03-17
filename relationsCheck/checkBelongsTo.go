package relationsCheck

import (
	"gormlint/common"
)

func IsBelongsTo(field common.Field, model common.Model, relatedModel common.Model) bool {
	foreignKey := field.Tags.GetParamOr("foreignKey", "Id")
	references := field.Tags.GetParamOr("references", relatedModel.Name+"Id")

	if !model.HasField(references) {
		return false
	}
	if !relatedModel.HasField(foreignKey) {
		return false
	}

	return true
}
