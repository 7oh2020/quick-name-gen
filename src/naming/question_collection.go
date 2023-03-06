package naming

import (
	"fmt"
	"strings"
)

// ユーザーへの質問の管理
type QuestionCollection struct {
	items []*Question
}

func NewQuestionCollection(items []*Question) *QuestionCollection {
	return &QuestionCollection{items}
}

// 質問を対話式に表示してユーザーから解答を入力してもらう
func (c *QuestionCollection) InputAnswers() error {
	for i := 0; i < len(c.items); i++ {
		var answer string
		fmt.Println(c.items[i].Value)
		if _, err := fmt.Scan(&answer); err != nil {
			return err
		}
		c.items[i].Answer = answer
	}
	return nil
}

// APIに渡す質問文を作成する
func (c QuestionCollection) CreateContent() string {
	// テンプレート内のIDを解答データで置き換える
	content := template
	for _, v := range c.items {
		content = strings.Replace(content, v.ID, v.Answer, 1)
	}
	return content
}
