package main

import (
	"github.com/kuzgoga/gormlint/nullSafetyCheck"
	"github.com/kuzgoga/gormlint/relationsCheck"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		nullSafetyCheck.NullSafetyAnalyzer,
		relationsCheck.RelationsAnalyzer,
	)
}
