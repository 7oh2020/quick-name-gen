package naming

// 質問データ
type Question struct {
	// 質問の種別(Who, What, Howのどれか)
	ID string

	// 質問
	Value string

	// 解答
	Answer string
}

func NewQuestion(id string, value string) *Question {
	return &Question{
		ID:    id,
		Value: value,
	}
}
