package todo_model

type (
	TodoRegistDTO struct {
		Content string `json:"content"`
		Writer  string `json:"writer"`
	}

	TodoModifyDTO struct {
		Content string `json:"content"`
	}
)
