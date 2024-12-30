package tests

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"gormlint/foreignKeyCheck"
	"testing"
)

func TestForeignKeyCheck(t *testing.T) {
	t.Parallel()
	analysistest.Run(t, analysistest.TestData(), foreignKeyCheck.ForeignKeyCheck, "foreign_key_check")
}
