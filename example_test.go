package tiffany_test

import (
	"fmt"
	"strings"

	"subosito.com/go/tiffany"
)

func ExampleRender() {
	out := &strings.Builder{}

	tiffany.Render(out, tiffany.Option{
		CanonicalURL: "subosito.com/go/gotenv",
		RepoURL:      "https://github.com/subosito/gotenv",
	})

	fmt.Println(out.String())
}
