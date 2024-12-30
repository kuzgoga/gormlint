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
