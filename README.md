Go RSS
=====

How to use this
===============

```go
import (
    "github.com/valerykalashnikov/gorss"
)

func feedHandler(w http.ResponseWriter, r *http.Request, posts) { 
    f := &gorss.Feed{
        Title:       "Title",
        Link:        "http://example.com",
        Description: "Your description",
    }
    for _, i := range(posts) {
        item := &gorss.Item {
                Title: i.Title,
                Link: i.Link,
                Description: i.Description,
            }
        f.AddItem(item)
    }
    out, err := f.Xml();
    if err != nil {
        http.Error(w,"", http.StatusInternalServerError)
    }
    fmt.Fprintf(w, out)
}
```