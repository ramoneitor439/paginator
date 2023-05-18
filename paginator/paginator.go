package paginator

import "errors"

type Page[T any] struct {
	Index       int
	HasPrevious bool
	HasNext     bool
	Items       []T
}

func (p *Page[T]) GetNext() int {
	if p.HasNext {
		return p.Index + 1
	}
	return -1
}

func (p *Page[T]) GetPrevious() int {
	if p.HasPrevious {
		return p.Index - 1
	}
	return -1
}

type Paginator[T any] struct {
	Items    []T
	PageSize int
}

func NewPaginator[T any](items []T, size int) *Paginator[T] {
	return &Paginator[T]{
		Items:    items,
		PageSize: size,
	}
}

func (pg *Paginator[T]) GetPage(index int) (*Page[T], error) {
	endIndex := index * pg.PageSize
	startIndex := endIndex - pg.PageSize

	if startIndex > len(pg.Items)-1 {
		return nil, errors.New("Page is above slice limits")
	}

	if len(pg.Items) < endIndex {
		endIndex = len(pg.Items)
	}

	result := Page[T]{
		Index:       index,
		HasNext:     false,
		HasPrevious: false,
		Items:       pg.Items[startIndex:endIndex],
	}
	if len(pg.Items) > endIndex+1 {
		result.HasNext = true
	}
	if index > 1 {
		result.HasPrevious = true
	}

	return &result, nil
}
