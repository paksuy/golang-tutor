# Golang Tutor

This repo contains support materials for learning go. basic-webservice3 is an example of implementation on using JSON as data



## Requirement

Familiar with Julien Schmidt's [httprouter](https://github.com/julienschmidt/httprouter) route and handler



## Tutorial

Working with JSON, we need to import go's JSON standard library

```go
import (
	"encoding/json"
)
```

### Sending Response in JSON

It boils down to

```go
func Handler(w http.ResponseWriter, r *http.Request, p httprouter.Params)
{
    data, err := json.Marshal(&structData)
    w.Header().Set("Content-Type", "application/json")
    w.Write(data)
}
```



## basic-webservice3

This branch contains basic web service implementation using JSON as data format.

### build

`make`

### test

`make test`



## Note

If you're using Windows, please download gnu Make for Windows



## Reference

* [HttpRouter](https://github.com/julienschmidt/httprouter): High Level Http Router