package common

import (
	"github.com/kuzgoga/fogg"
	"go/ast"
	"go/token"
)

type Field struct {
	Name    string
	Type    ast.Expr
	Tag     *string
	Pos     token.Pos
	Comment string
	Tags    fogg.Tag
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
