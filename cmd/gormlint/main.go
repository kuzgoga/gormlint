package main

import (
	"golang.org/x/tools/go/analysis/multichecker"
	"gormlint/nullSafetyCheck"
	"gormlint/relationsCheck"
)

func main() {
	multichecker.Main(
		nullSafetyCheck.NullSafetyAnalyzer,
		relationsCheck.RelationsAnalyzer,
	)
}
