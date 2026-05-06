package formatter

func Format(text string) string {
	text = ConvertStyle(text)
	return text
}
