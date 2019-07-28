# tiffany

[![Build Status](https://api.cirrus-ci.com/github/subosito/tiffany.svg)](https://cirrus-ci.com/github/subosito/tiffany)
[![Coverage Status](https://badgen.net/codecov/c/github/subosito/tiffany)](https://codecov.io/gh/subosito/tiffany)
[![GoDoc](https://godoc.org/github.com/subosito/tiffany?status.svg)](https://godoc.org/subosito.com/go/tiffany)

Go vanity URL.

## Usage

```go
import "subosito.com/go/tiffany"
```

For example, you can create a server for your vanity import path:

```go
import (
	"net/http"

	"subosito.com/go/tiffany"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Cache-Control", "public, max-age=300")

    tiffany.Render(w, tiffany.Option{
        CanonicalURL: "subosito.com/go/gotenv",
        RepoURL:      "https://github.com/subosito/gotenv",
    })
}

http.Handle("/go/gotenv", Handler)
http.ListenAndServe(":8080", nil)
```
