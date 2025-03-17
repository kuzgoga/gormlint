package tests

import (
	"github.com/kuzgoga/gormlint/nullSafetyCheck"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestNullSafety(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), nullSafetyCheck.NullSafetyAnalyzer, "null_safety")
}
