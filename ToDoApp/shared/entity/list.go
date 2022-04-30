package entity

import . "github.com/google/uuid"

type List struct {
	id   string
	name string
}

type ListId struct {
	id UUID
}

func (l ListId) ToString() string {
	return l.id.String()
}

type ListName struct {
	name string
}

func (l ListName) Name() string {
	return l.name
}
