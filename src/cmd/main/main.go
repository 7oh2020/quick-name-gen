package main

import (
	"fmt"
	"os"
	"quick-name-gen/src/chat"
	"quick-name-gen/src/naming"
	"time"
)

func main() {
	// 環境変数からAPIキーを取得する
	secret, ok := os.LookupEnv("OPEN_AI_SECRET")
	if !ok {
		panic("open-api-secret is empty")
	}

	// リソース節約のためにタイムアウトを設定する
	timeout := 60 * time.Second

	// トークン節約のために応答の最大トークンを設定する
	maxTokens := 1500

	// チャットに使用するモデルのID
	modelID := "gpt-3.5-turbo"

	// ユーザーへの3つの質問を作成する
	items := []*naming.Question{
		naming.NewQuestion("[who]", "そのアプリは誰のためのアプリですか？"),
		naming.NewQuestion("[what]", "そのアプリは人々に何を与えますか？"),
		naming.NewQuestion("[how]", "それはどんな方法で与えられますか？"),
	}
	qc := naming.NewQuestionCollection(items)

	// 対話式にユーザーの解答を取得する
	if err := qc.InputAnswers(); err != nil {
		panic(err)
	}

	// 解答データからAPIへの質問文を作成する
	content := qc.CreateContent()
	fmt.Println("...")

	// APIにメッセージを送信する
	c := chat.NewChatCompletions(modelID, secret, maxTokens, timeout)
	res, err := c.AskOneQuestion(content)
	if err != nil {
		panic(err)
	}

	fmt.Printf("In %d / Out %d / Total %d tokens\n", res.Usage.PromptTokens, res.Usage.CompletionTokens, res.Usage.TotalTokens)
	for _, v := range res.Choices {
		fmt.Println(v.Message.Content)
	}
}
