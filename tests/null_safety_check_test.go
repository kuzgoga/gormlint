package tests

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"gormlint/nullSafetyCheck"
	"testing"
)

func TestNullSafety(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), nullSafetyCheck.NullSafetyAnalyzer, "null_safety")
}
