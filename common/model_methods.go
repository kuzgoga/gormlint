package common

func (model *Model) HasPrimaryKey() bool {
	for _, field := range model.Fields {
		if field.Tags.HasParam("primaryKey") || field.Tags.HasOption("primaryKey") {
			return true
		}
	}
	return false
}
