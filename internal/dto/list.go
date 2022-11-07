package dto

import (
	"bytes"
)

const defaultListLimit = 50

type Order string

const (
	ASC  Order = "asc"
	DESC Order = "desc"
)

type ListFilter struct {
	Limit  uint
	Offset uint
	Order  string
}

func NewListFilter(limit, offset uint, order Order, sortBy string) ListFilter {
	if limit == 0 || limit > defaultListLimit {
		limit = defaultListLimit
	}

	if order != ASC && order != DESC {
		order = DESC
	}

	var buffer bytes.Buffer
	if len(sortBy) != 0 {
		buffer.WriteString(sortBy)
	} else {
		buffer.WriteString("id")
	}
	buffer.WriteString(" ")
	buffer.WriteString(string(order))

	return ListFilter{
		Limit:  limit,
		Offset: offset,
		Order:  buffer.String(),
	}
}
