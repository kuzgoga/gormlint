package relationsCheck

import "gormlint/common"

func IsHasOne(field common.Field, model common.Model, relatedModel common.Model) bool {
	foreignKey := field.Tags.GetParamOr("foreignKey", model.Name+"Id")
	references := field.Tags.GetParamOr("references", "Id")

	if !relatedModel.HasField(foreignKey) {
		return false
	}

	if !model.HasField(references) {
		return false
	}

	return true
}
