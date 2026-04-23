package formatter

import (
	"fmt"
	"regexp"
	"strings"
)

type styleRule struct {
	re          *regexp.Regexp
	replacement string
}

func r(pattern, replacement string) styleRule {
	return styleRule{regexp.MustCompile(pattern), replacement}
}

var styleRules = []styleRule{
	r(`ではありませんでした`, `ではなかった`),
	r(`ありませんでした`, `なかった`),
	r(`できませんでした`, `できなかった`),
	r(`りませんでした`, `らなかった`),
	r(`きませんでした`, `かなかった`),
	r(`ぎませんでした`, `がなかった`),
	r(`びませんでした`, `ばなかった`),
	r(`みませんでした`, `まなかった`),
	r(`にませんでした`, `ななかった`),
	r(`ちませんでした`, `たなかった`),
	r(`しませんでした`, `しなかった`),
	r(`ませんでした`, `なかった`),

	r(`ではありません`, `ではない`),
	r(`ではないです`, `ではない`),
	r(`ありません`, `ない`),
	r(`できません`, `できない`),
	r(`りません`, `らない`),
	r(`きません`, `かない`),
	r(`ぎません`, `がない`),
	r(`びません`, `ばない`),
	r(`みません`, `まない`),
	r(`にません`, `なない`),
	r(`ちません`, `たない`),
	r(`しません`, `しない`),
	r(`ません`, `ない`),

	r(`ましょう`, `よう`),
	r(`ました`, `た`),

	r(`でした`, `であった`),
	r(`でしょう`, `であろう`),
	r(`ですが`, `であるが`),
	r(`ですけれど`, `であるけれど`),
	r(`ですけど`, `であるけど`),
	r(`ですから`, `であるから`),
	r(`ですので`, `であるので`),
	r(`ですし`, `であるし`),
	r(`ですね`, `である`),
	r(`ですよ`, `である`),
	r(`です`, `である`),

	r(`ています`, `ている`),
	r(`てきます`, `てくる`),
	r(`てみます`, `てみる`),
	r(`しておきます`, `する`),
	r(`ておきます`, `ておく`),

	r(`できます`, `できる`),
	r(`起きます`, `起きる`),
	r(`着ます`, `着る`),

	r(`します`, `する`),

	r(`きます`, `く`),
	r(`ぎます`, `ぐ`),
	r(`びます`, `ぶ`),
	r(`みます`, `む`),
	r(`にます`, `ぬ`),
	r(`ります`, `る`),
	r(`ちます`, `つ`),

	{regexp.MustCompile(`(\p{Han})います`), `${1}う`},
	r(`います`, `いる`),

	{regexp.MustCompile(`([\p{Han}ぁ-んァ-ヴ])ます`), `${1}る`},
}

var (
	codeFenceRe  = regexp.MustCompile("(?s)```.*?```")
	inlineCodeRe = regexp.MustCompile("`[^`\n]+`")
)

func ConvertStyle(text string) string {
	var blocks []string

	protect := func(s string) string {
		id := len(blocks)
		blocks = append(blocks, s)
		return fmt.Sprintf("\x00%d\x00", id)
	}

	text = codeFenceRe.ReplaceAllStringFunc(text, protect)
	text = inlineCodeRe.ReplaceAllStringFunc(text, protect)

	for _, rule := range styleRules {
		text = rule.re.ReplaceAllString(text, rule.replacement)
	}

	for i, block := range blocks {
		text = strings.ReplaceAll(text, fmt.Sprintf("\x00%d\x00", i), block)
	}
	return text
}
