# ai-formatter

AIが生成しがちな日本語テキストをドキュメント体（だ・である調）に変換するライブラリ

- ですます調 → だ・である調への変換
- Markdown の太字マーカー（`**text**` → `text`）の除去

## CLIの使い方

```sh
go run ./cmd/ < input.txt
```

## wasmビルド

```sh
GOOS=js GOARCH=wasm go build -o formatter.wasm ./cmd/wasm/
```

`wasm_exec.js` はGoの標準ライブラリから取得する。

```sh
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" .
```

### ブラウザでの使い方

`formatter.wasm` と `wasm_exec.js` を配信し、以下のように読み込む。
wasmの読み込み後、`window.aiFormat(text)` が利用できる:

```html
<script src="wasm_exec.js"></script>
<script>
  const go = new Go();
  WebAssembly.instantiateStreaming(fetch("formatter.wasm"), go.importObject).then(result => {
    go.run(result.instance);
    console.log(window.aiFormat("これは重要です。"));
    // → "これは重要である。"
  });
</script>
```

## License

MIT
