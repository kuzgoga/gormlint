package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"gormlint/nullSafetyCheck"
)

func main() {
	singlechecker.Main(nullSafetyCheck.NullSafetyAnalyzer)
}
