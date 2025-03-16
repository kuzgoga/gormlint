package common

import (
	"github.com/kuzgoga/fogg"
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

func FindModelParam(paramName string, model Model) *fogg.TagParam {
	for _, field := range model.Fields {
		if field.Tags.HasParam(paramName) {
			return field.Tags.GetParam(paramName)
		}
	}
	return nil
}

func FindReferencesInM2M(m2mReference Field, relatedModel Model) *Field {
	/* Find `references` field in m2m relation */
	referencesTag := m2mReference.Tags.GetParam("references")
	if referencesTag != nil {
		return GetModelField(&relatedModel, referencesTag.Value)
	} else {
		for _, field := range relatedModel.Fields {
			if field.Tags.HasOption("primaryKey") {
				return &field
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
		m2mRelation := field.Tags.GetParam("many2many")
		if m2mRelation != nil {
			if m2mRelation.Value == relationName {
				return &field
			}
		}
	}
	return nil
}

//func findForeignKey()
