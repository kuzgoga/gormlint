package tests

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"gormlint/nullSafetyCheck"
	"testing"
)

func TestNullSafety(t *testing.T) {
	t.Parallel()
	analysistest.Run(t, analysistest.TestData(), nullSafetyCheck.NullSafetyAnalyzer, "null_safety")
}
