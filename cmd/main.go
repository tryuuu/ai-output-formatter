package main

import (
	"fmt"
	"io"
	"os"

	"github.com/tryuuu/ai-formatter/internal/formatter"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Print(formatter.Format(string(input)))
}
