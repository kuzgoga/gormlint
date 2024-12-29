package common

import "strings"

func NormalizeStructTags(tags string) string {
	// todo: process case with check with ';' literal
	tagWithoutQuotes := tags[1 : len(tags)-1]
	tagWithoutSemicolons := strings.ReplaceAll(tagWithoutQuotes, ";", ",")
	return tagWithoutSemicolons
}
