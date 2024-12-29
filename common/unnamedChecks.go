package common

import (
	"errors"
	"fmt"
	"go/ast"
)

func CheckUnnamedModel(typeSpec ast.TypeSpec) error {
	if typeSpec.Name == nil {
		return errors.New(fmt.Sprintf("Unnamed model\n"))
	}
	return nil
}

func CheckUnnamedField(structName string, field ast.Field) error {
	if len(field.Names) == 0 {
		return errors.New(fmt.Sprintf("Struct \"%s\" has unnamed field", structName))
	}
	return nil
}
