package formatter

func Format(text string) string {
	text = RemoveBold(text)
	text = ConvertStyle(text)
	return text
}
