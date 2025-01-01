package common

import (
	"go/ast"
	"go/token"
	"strings"
)

type Field struct {
	Name    string
	Type    ast.Expr
	Tags    *string
	Options []string // contains options like "autoCreateTime" or "null"
	Params  []string // contains params like "foreignKey:CustomerId" or "constrain:OnDelete:Cascade"
	Pos     token.Pos
	Comment string
}

type Model struct {
	Name    string
	Fields  map[string]Field
	Pos     token.Pos
	Comment string
}

type Param struct {
	Name  string
	Value string
}

func (model *Model) GetParam(name string) *Param {
	for _, field := range model.Fields {
		for _, param := range field.Params {
			pair := strings.SplitN(param, ":", 2)
			if len(pair) != 2 {
				return nil
			}
			if strings.ToLower(pair[0]) == strings.ToLower(name) {
				return &Param{
					Name:  pair[0],
					Value: pair[1],
				}
			}
		}
	}
	return nil
}

func (model *Model) HasParam(name string) bool {
	return model.GetParam(name) != nil
}

func (field *Field) HasParam(name string) bool {
	return field.GetParam(name) != nil
}

func (field *Field) GetParam(name string) *Param {
	for _, param := range field.Params {
		pair := strings.SplitN(param, ":", 2)
		if len(pair) != 2 {
			return nil
		}
		if strings.ToLower(pair[0]) == strings.ToLower(name) {
			return &Param{
				Name:  pair[0],
				Value: pair[1],
			}
		}
	}
	return nil
}
