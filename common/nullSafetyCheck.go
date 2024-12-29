package common

import (
	"errors"
	"fmt"
	"github.com/fatih/structtag"
	"go/ast"
)

func isPointerType(typeExpr ast.Expr) bool {
	isPointer := false
	if _, ok := typeExpr.(*ast.StarExpr); ok {
		isPointer = true
	}
	return isPointer
}

func isGormValueNullable(tags *structtag.Tags) (*bool, error) {
	gormTag, err := tags.Get("gorm")
	if gormTag == nil {
		return nil, nil
	}

	gormTag.Options = append([]string{gormTag.Name}, gormTag.Options...)

	if err != nil {
		return nil, nil
	}

	nullTagExist := gormTag.HasOption("null")
	notNullTagExist := gormTag.HasOption("not null")

	if nullTagExist == notNullTagExist && nullTagExist {
		return nil, errors.New(`tags "null" and "not null" are specified at one field`)
	}

	if nullTagExist {
		return PointerOf(true), nil
	} else if notNullTagExist {
		return PointerOf(false), nil
	} else {
		return PointerOf(false), nil
	}
}

func CheckFieldNullConsistency(field ast.Field, structName string, structTags string) error {
	tags, err := structtag.Parse(structTags)
	if err != nil {
		return errors.New(fmt.Sprintf("Invalid structure tag: %s", err))

	}
	if tags == nil {
		return nil
	}

	isFieldNullable := isPointerType(field.Type)
	isColumnNullable, err := isGormValueNullable(tags)

	if err != nil {
		return errors.New(fmt.Sprintf("Null safety error: %s", err))
	}
	if isColumnNullable == nil {
		return nil
	}

	if isFieldNullable != *isColumnNullable {
		return errors.New(fmt.Sprintf("Null safety error in \"%s\" model, field \"%s\": column nullable policy doesn't match to tag nullable policy", structName, field.Names[0].Name))
	}
	return nil
}
