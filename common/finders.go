package common

import (
	"go/ast"
	"strings"
)

func GetModelField(model *Model, fieldName string) *Field {
	field, fieldExists := model.Fields[fieldName]
	if fieldExists {
		return &field
	} else {
		return nil
	}
}

func GetModelFromType(modelType ast.Expr, models map[string]Model) *Model {
	baseType := ResolveBaseType(modelType)
	if baseType != nil {
		return GetRelatedModel(*baseType, models)
	} else {
		return nil
	}
}

func GetRelatedModel(modelName string, models map[string]Model) *Model {
	model, modelExists := models[modelName]
	if modelExists {
		return &model
	} else {
		return nil
	}
}

func FindParamValue(paramName string, params []string) *string {
	for _, rawParam := range params {
		pair := strings.Split(rawParam, ":")
		if len(pair) < 2 {
			return nil
		}
		if strings.ToLower(pair[0]) == strings.ToLower(paramName) {
			return &pair[1]
		}
	}
	return nil
}

func FindModelParam(paramName string, model Model) *Param {
	for _, field := range model.Fields {
		for _, param := range field.Params {
			pair := strings.Split(param, ":")
			if len(pair) < 2 {
				return nil
			}
			if strings.ToLower(pair[0]) == strings.ToLower(paramName) {
				return &Param{
					Name:  pair[0],
					Value: pair[1],
				}
			}
		}
	}
	return nil
}

func FindReferencesInM2M(m2mReference Field, relatedModel Model) *Field {
	/* Find `references` field in m2m relation */
	referencesTagValue := FindParamValue("references", m2mReference.Params)
	if referencesTagValue != nil {
		return GetModelField(&relatedModel, *referencesTagValue)
	} else {
		for _, field := range relatedModel.Fields {
			for _, opt := range field.Options {
				if opt == "primaryKey" {
					return &field
				}
			}
		}
		for _, field := range relatedModel.Fields {
			if strings.ToLower(field.Name) == "id" {
				return &field
			}
		}
		return nil
	}
}

func FindBackReferenceInM2M(relationName string, relatedModel Model) *Field {
	for _, field := range relatedModel.Fields {
		m2mRelation := field.GetParam("many2many")
		if m2mRelation != nil {
			if m2mRelation.Value == relationName {
				return &field
			}
		}
	}
	return nil
}

//func findForeignKey()
