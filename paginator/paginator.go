package paginator

import "errors"

type Page struct {
	Index       int
	HasPrevious bool
	HasNext     bool
	Items       []any
}

func (p *Page) GetNext() (int, error) {
	if p.HasNext {
		return p.Index + 1, nil
	}
	return -1, errors.New("Index out of paginator limits")
}

func (p *Page) GetPrevious() (int, error) {
	if p.HasPrevious {
		return p.Index - 1, nil
	}
	return -1, errors.New("Index out of paginator limits")
}

type Paginator struct {
	Items    []any
	PageSize int
}

func NewPaginator(items []any, size int) (*Paginator, error) {
	return &Paginator{
		Items:    items,
		PageSize: size,
	}, nil
}

func (pg *Paginator) GetPage(index int) (*Page, error) {
	endIndex := index * pg.PageSize
	startIndex := endIndex - pg.PageSize

	if startIndex > len(pg.Items)-1 {
		return nil, errors.New("Page is above slice limits")
	}

	if len(pg.Items) < endIndex {
		endIndex = len(pg.Items)
	}

	result := Page{
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