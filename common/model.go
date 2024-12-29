package common

import (
	"go/ast"
	"go/token"
)

type Field struct {
	Name     string
	Type     ast.Expr
	Tags     *string
	Options  []string // contains options like "autoCreateTime" or "null"
	Params   []string // contains params like "foreignKey:CustomerId" or "constrain:OnDelete:Cascade"
	Position token.Pos
	Comment  string
}

type Model struct {
	Name     string
	Fields   map[string]Field
	Position token.Pos
	Comment  string
}
