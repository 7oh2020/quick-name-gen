package main

import (
	"fmt"
	"os"
	"quick-name-gen/src/chat"
	"time"
)

func main() {
	// コマンドライン引数から質問テキストを取得する
	if len(os.Args) < 2 {
		panic("too few arguments")
	}
	content := os.Args[1]

	// 環境変数からAPIキーを取得する
	secret, ok := os.LookupEnv("OPEN_AI_SECRET")
	if !ok {
		panic("open-api-secret is empty")
	}

	// リソース節約のためにタイムアウトを設定する
	timeout := 15 * time.Second

	// トークン節約のために応答の最大トークンを設定する
	maxTokens := 500

	// チャットに使用するモデルのID
	modelID := "gpt-3.5-turbo"

	c := chat.NewChatCompletions(modelID, secret, maxTokens, timeout)
	res, err := c.AskOneQuestion(content)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("In %d / Out %d / Total %d tokens\n", res.Usage.PromptTokens, res.Usage.CompletionTokens, res.Usage.TotalTokens)
	for _, v := range res.Choices {
		fmt.Printf("[%s]: %s\n", v.Message.Role, v.Message.Content)
	}

}
