package relationsCheck

import (
	"gormlint/common"
)

func IsBelongsTo(field common.Field, model common.Model, relatedModel common.Model) bool {
	references := field.Tags.GetParamOr("references", "Id")
	foreignKey := field.Tags.GetParamOr("foreignKey", field.Name+"Id")

	if !model.HasField(foreignKey) {
		return false
	}
	if !relatedModel.HasField(references) {
		return false
	}

	return true
}
