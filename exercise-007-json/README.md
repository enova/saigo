## Description

We are going to learn how to encode and decode JSON in Go. You've been given some code that reads data from a file and decodes the JSON into a struct. Your job is to complete the implementation.

## Comprehension Tasks

The json package implements encoding and decoding of JSON objects.

Check out this code:

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Element struct {
	Name         string
	Symbol       string
	AtomicNumber int
	AtomicWeight float64
	Category     string
	Group        int
	Period       int
}

func main() {

	e := Element{
		Name:         "Gold",
		Symbol:       "Au",
		AtomicNumber: 79,
		AtomicWeight: 196.966,
		Category:     "transition metal",
		Group:        11,
		Period:       6,
	}

	data, err := json.Marshal(&e)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
```

This code is from [exhibit-a](). Run the code in the console:

```
exercise-007-json $ go run exhibit-a/app.go
```

For prettier output, there is a method in the `json` package called `MarshalIndent()`.
Try using this method in place of `json.Marshal` in the code and see the output.

Read the code in each of the following directories and explain them back to an instructor:

1. [exhibit-a]() Marshal
1. [exhibit-b]() Tags
1. [exhibit-c]() Unmarshal
1. [exhibit-d]() JSON API!

## Engineering Tasks

In [exhibit-d](), notice how much of the data in the JSON file is discarded when populating phone data:

```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Phone ...
type Phone struct {
    Name string `json:"name"`
}

var allPhones []Phone

func setup() {
    data, err := ioutil.ReadFile("phones.json")
    if err != nil {
        fmt.Println("Error reading phones.json")
        os.Exit(1)
      }

    err = json.Unmarshal(data, &allPhones)
    if err != nil {
        fmt.Println("Error in unmarshalling phones")
    }
}

func phones(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(allPhones)
}

func main() {
    setup()
    http.HandleFunc("/phones", phones)
    http.ListenAndServe(":8080", nil)
}
```

Complete the following tasks:

1. Read the other attributes from the file and store them in the `Phone` slice.
2. Create a database table for phones. When the server is started, if the table is empty, seed the table with the phones from the json file. The /phones endpoint should then return whatever records are found in the phones table.
