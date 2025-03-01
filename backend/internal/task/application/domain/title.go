package domain

type Title struct {
	value string
}

func NewTitle(title string) (Title, error) {
	return Title{value: title}, nil
}

func (t *Title) Value() string {
	return t.value
}