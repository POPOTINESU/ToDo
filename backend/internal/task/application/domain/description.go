package domain

const MAX_DESCRIPTION_LENGTH = 255

type Description struct {
	value string
}

func NewDescription(description string) (Description, error) {
	return Description{
		value: description,
	}, nil
}

func (v Description) Value() string {
	return v.value
}