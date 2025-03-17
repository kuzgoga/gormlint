package tests

import (
	"github.com/kuzgoga/gormlint/relationsCheck"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestRelationsCheck(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), relationsCheck.RelationsAnalyzer, "relations_check")
}
