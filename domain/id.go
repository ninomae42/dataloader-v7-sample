package domain

type ID string

func NewID() ID {
	// NOTE: mock value
	return ID("uuid")
}

func (id ID) String() string {
	return string(id)
}
