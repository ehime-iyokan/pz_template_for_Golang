package main

import (
	"bufio"
	"os"
)

func MakeAnswer() (string, error) {
	// 入力用の Scanner 定義処理：(解答作成時)
	f, err := os.Open("input.txt")
	if err != nil {
		return "", err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)

	// 入力用の Scanner 定義処理：(解答提出時)
	// sc := bufio.NewScanner(os.Stdin)

	// 以下、解答を記載する
	ans := ""

	return ans, nil
}
