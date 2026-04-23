package formatter

import "regexp"

var boldRe = regexp.MustCompile(`\*\*(.+?)\*\*`)

// **を削除
func RemoveBold(text string) string {
	return boldRe.ReplaceAllString(text, "$1")
}
