package relationsCheck

import (
	"github.com/kuzgoga/gormlint/common"
	"golang.org/x/tools/go/analysis"
	"strings"
)

func CheckCascadeDelete(pass *analysis.Pass, field common.Field) bool {
	if !field.Tags.HasParam("constraint") {
		pass.Reportf(field.Pos, "field %s should have a delete constraint", field.Name)
		return true
	}
	constraintValue := field.Tags.GetParam("constraint").Value
	pair := strings.Split(constraintValue, ":")
	trigger, value := pair[0], pair[1]
	if strings.ToLower(trigger) == "OnDelete" && strings.ToUpper(value) != "CASCADE" {
		pass.Reportf(field.Pos, "field have invalid constraint on `OnDelete trigger`")
		return true
	}

	if field.Tags.HasParam("OnDelete") {
		if strings.ToUpper(field.Tags.GetParam("OnDelete").Value) != "CASCADE" {
			pass.Reportf(field.Pos, "field have invalid constraint on `OnDelete` trigger")
			return true
		}
	}
	return false
}
