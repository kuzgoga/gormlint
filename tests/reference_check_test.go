package tests

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"gormlint/referencesCheck"
	"testing"
)

func TestReferenceCheck(t *testing.T) {
	t.Parallel()
	analysistest.Run(t, analysistest.TestData(), referencesCheck.ReferenceAnalyzer, "references_check")
}
