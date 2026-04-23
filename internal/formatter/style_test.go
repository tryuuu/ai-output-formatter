package formatter

import "testing"

func TestConvertStyle(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "ります→る (仕様例)",
			input: "図では、セグメント1、2、3という3つのファイルがあります。同じキー（例えば handful）が複数のセグメントに存在しますが、それぞれ値が異なります。",
			want:  "図では、セグメント1、2、3という3つのファイルがある。同じキー（例えば handful）が複数のセグメントに存在するが、それぞれ値が異なる。",
		},
		{
			name:  "です→である",
			input: "これは重要です。",
			want:  "これは重要である。",
		},
		{
			name:  "ですが→であるが",
			input: "効果的ですが、注意が必要です。",
			want:  "効果的であるが、注意が必要である。",
		},
		{
			name:  "でした→であった",
			input: "結果は良好でした。",
			want:  "結果は良好であった。",
		},
		{
			name:  "でしょう→であろう",
			input: "原因はこれでしょう。",
			want:  "原因はこれであろう。",
		},
		{
			name:  "ません→ない (ru-verb catch-all)",
			input: "食べません。",
			want:  "食べない。",
		},
		{
			name:  "りません→らない",
			input: "変わりません。",
			want:  "変わらない。",
		},
		{
			name:  "きません→かない",
			input: "書きません。",
			want:  "書かない。",
		},
		{
			name:  "ありません→ない",
			input: "問題はありません。",
			want:  "問題はない。",
		},
		{
			name:  "ではありません→ではない",
			input: "これは問題ではありません。",
			want:  "これは問題ではない。",
		},
		{
			name:  "できません→できない",
			input: "実行できません。",
			want:  "実行できない。",
		},
		{
			name:  "ませんでした→なかった",
			input: "食べませんでした。",
			want:  "食べなかった。",
		},
		{
			name:  "りませんでした→らなかった",
			input: "変わりませんでした。",
			want:  "変わらなかった。",
		},
		{
			name:  "ました→た",
			input: "処理しました。",
			want:  "処理した。",
		},
		{
			name:  "ましょう→よう",
			input: "確認しましょう。",
			want:  "確認しよう。",
		},
		{
			name:  "します→する",
			input: "存在しますが、値が異なります。",
			want:  "存在するが、値が異なる。",
		},
		{
			name:  "ています→ている",
			input: "処理しています。",
			want:  "処理している。",
		},
		{
			name:  "しておきます→する",
			input: "記録しておきます。",
			want:  "記録する。",
		},
		{
			name:  "できます→できる (ru-verb例外)",
			input: "実行できます。",
			want:  "実行できる。",
		},
		{
			name:  "います→いる (単独)",
			input: "ユーザーがいます。",
			want:  "ユーザーがいる。",
		},
		{
			name:  "います→う (漢字先行)",
			input: "AをBと言います。",
			want:  "AをBと言う。",
		},
		{
			name:  "一段動詞フォールバック",
			input: "データを忘れます。",
			want:  "データを忘れる。",
		},
		{
			name:  "コードブロック保護",
			input: "以下を実行します。\n```\nfoo bar\n```\n完了します。",
			want:  "以下を実行する。\n```\nfoo bar\n```\n完了する。",
		},
		{
			name:  "インラインコード保護",
			input: "`foo`を使います。",
			want:  "`foo`を使う。",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertStyle(tt.input)
			if got != tt.want {
				t.Errorf("\ninput: %q\n  got: %q\n want: %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestRemoveBold(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "単純なbold除去",
			input: `**テキスト**`,
			want:  `テキスト`,
		},
		{
			name:  "文中のbold除去",
			input: `**「キーが同じなら最新を採用する」**というルールでマージします。`,
			want:  `「キーが同じなら最新を採用する」というルールでマージします。`,
		},
		{
			name:  "複数bold除去",
			input: `**A**と**B**があります。`,
			want:  `AとBがあります。`,
		},
		{
			name:  "boldなし (変化なし)",
			input: `通常のテキストです。`,
			want:  `通常のテキストです。`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveBold(tt.input)
			if got != tt.want {
				t.Errorf("\ninput: %q\n  got: %q\n want: %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	input := `最新の値だけを残す: **「キーが同じなら、最も新しいセグメントにある値を採用する」**というルールでマージします。`
	want := `最新の値だけを残す: 「キーが同じなら、最も新しいセグメントにある値を採用する」というルールでマージする。`

	got := Format(input)
	if got != want {
		t.Errorf("\ninput: %q\n  got: %q\n want: %q", input, got, want)
	}
}
