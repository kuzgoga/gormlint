package relationsCheck

import (
	"fmt"
	"github.com/kuzgoga/gormlint/common"
	"go/types"
	"golang.org/x/tools/go/analysis"
)

func findBackReferenceInOneToMany(model common.Model, relatedModel common.Model) *common.Field {
	for _, field := range relatedModel.Fields {
		if !common.IsSlice(field.Type) {
			continue
		}
		if field.Tags.HasParam("many2many") {
			continue
		}
		baseType := common.ResolveBaseType(field.Type)
		if baseType == nil {
			continue
		}
		if *baseType == model.Name {
			return &field
		}
	}
	return nil
}

func isOneToMany(pass *analysis.Pass, model common.Model, relatedModel common.Model) bool {
	backReference := findBackReferenceInOneToMany(model, relatedModel)
	if backReference == nil {
		return false
	}
	fmt.Println("Found back reference")
	fmt.Printf("Backref type: %s\n", types.ExprString(backReference.Type))
	fmt.Printf("Model: %s\n", model.Name)
	fmt.Printf("Related model: %s\n", relatedModel.Name)
	if checkManyToOne(pass, *backReference, relatedModel, model) {
		return false
	}
	return true
}
