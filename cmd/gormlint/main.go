package main

import (
	"golang.org/x/tools/go/analysis/multichecker"
	"gormlint/foreignKeyCheck"
	"gormlint/nullSafetyCheck"
	"gormlint/referencesCheck"
)

func main() {
	multichecker.Main(
		nullSafetyCheck.NullSafetyAnalyzer,
		referencesCheck.ReferenceAnalyzer,
		foreignKeyCheck.ForeignKeyCheck,
	)
}
