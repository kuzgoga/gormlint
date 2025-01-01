package tests

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"gormlint/relationsCheck"
	"testing"
)

func TestRelationsCheck(t *testing.T) {
	t.Parallel()
	analysistest.Run(t, analysistest.TestData(), relationsCheck.RelationsAnalyzer, "relations_check")
}
