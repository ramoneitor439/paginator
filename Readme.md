### Requirements

```
    Go v1.18 or higher
```

## Instalation

Write this command in your shell, placed in your project folder

``` bash
    go get github.com/ramoneitor439/paginator
```

## Attributes



## Paginator

---
#### Items:
```
    Returns the list of items the paginator will paginate
```
#### PageSize:
```
    Returns the ammount of elements that every page will contain
```
#### Length:
```
    Returns the ammount of pages in the paginator
```
#### FirstIndex:
```
    Returns the index of the first page in the paginator
```
#### LastIndex:
```
    Returns the last page in the paginator
```
## Page

---
#### Index:
```
    Returns the index of the current page
```
#### HasNext:
```
    Returns true if exists another page after current
```
#### HasPrevious:
```
    Returns true if exists another page before current
```
#### Items:
```
    Returns the list of items in the current page
```
---

## Methods
```go
paginator.NewPaginator[type](data []any, pageSize int) *Paginator[type] // Create the paginator 
```
## Paginator

---
```go
GetPage(index int) *Page[type] // Creates a index based page 
```

## Page

---
```go
GetNext() int // returns index of next page; returns -1 if does not exist
```
```go
GetPrevious() int // returns index of previous page; returns -1 if does not exist
```



## Usage

``` go
    package example

    import "github.com/ramoneitor439/paginator"
    
    func Foo(pageNumber int, pageSize int) *Page[string] { // Define the type of data wich will contain the pages
        data := []string{"Element1", "Element2", "Element3", "Element4", "Element5"} // Get all the data you want to paginate
        
        pg := paginator.NewPaginator[string](data, pageSize) // Create a paginator according to the data type and the page size
        page := pg.GetPage(pageNumber) // Get the page according to the page number
        
        // Methods
        
        // page.HasNext -- returns true if there is another page next to it available
        
        // page.HasPrevious -- returns true if there is another page previous to it available
        
        // page.Index -- returns page index
        
        // page.GetNext() -- returns next page index; if there is no next page returns -1
        
        // page.GetPrevious() -- returns previous page index, if there is no previous page returns -1
        
        // page.Items -- returns an slice of the selected type with all page elements
        
        return page // return page for other specific usages  
    }
```