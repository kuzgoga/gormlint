package foreignKeyCheck

import (
	"golang.org/x/tools/go/analysis"
	"gormlint/common"
	"strings"
)

// ForeignKeyCheck todo: add URL for foreign key analyzer rules
var ForeignKeyCheck = &analysis.Analyzer{
	Name: "GormForeignKeyCheck",
	Doc:  "Check foreign key in gorm model struct tag",
	Run:  run,
}

var models map[string]common.Model

func run(pass *analysis.Pass) (any, error) {
	models = make(map[string]common.Model)
	common.ParseModels(pass, &models)

	for _, model := range models {
		for _, field := range model.Fields {
			for _, param := range field.Params {
				pair := strings.Split(param, ":")
				paramName := pair[0]
				paramValue := pair[1]
				if paramName == "foreignKey" {
					foreignKey, fieldExist := model.Fields[paramValue]

					if !fieldExist {
						pass.Reportf(field.Pos, "Foreign key \"%s\" mentioned in tag at field \"%s\" doesn't exist in model \"%s\"", paramValue, field.Name, model.Name)
					} else {
						foreignKeyType := common.ResolveBaseType(foreignKey.Type)
						if foreignKeyType == nil {
							pass.Reportf(foreignKey.Pos, "Failed to resolve type of foreign key field \"%s\": %s", field.Name, foreignKey.Type)
						} else {
							// TODO: handle all int types
							if *foreignKeyType != "uint" && *foreignKeyType != "int" {
								pass.Reportf(foreignKey.Pos, "Foreign key should have type like int, not \"%s\"", foreignKey.Type)
							}
						}
					}
				}
			}
		}
	}
	return nil, nil
}
