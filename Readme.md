## Instalation

Write this command in your shell
``` bash
    go get github.com/ramoneitor439/paginator
```

## Usage

``` go
    import "github.com/ramoneitor439/paginator"
    
    func GetContentPaginator(pageNumber int, pageSize int) interface{} {
        data := GetAllData() // Get all the data you want to paginate
        
        var dataAny []interface{} // Create an interface slice to save results         
        for _, obj := range data {
            dataAny = append(dataAny, obj)
        }
        
        pg := paginator.NewPaginator(dataAny, pageSize) // Create a paginator according to the data and the page size
        page := pg.GetPage(pageNumber) // Get the page according to the page number
        
        // Methods
        
        page.HasNext // returns true if there is another page next to it available
        
        page.HasPrevious // returns true if there is another page previous to it available
        
        page.Index // returns page index
        
        page.GetNext() // returns next page index; if there is no next page returns -1
        
        page.GetPrevious() // returns previous page index, if there is no previous page returns -1
        
        page.Items // returns an []any with all page elements
        
    }
```