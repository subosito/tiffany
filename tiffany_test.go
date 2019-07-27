package tiffany_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subosito/tiffany"
)

func TestRender(t *testing.T) {
	opt := tiffany.Option{
		CanonicalURL: "subosito.com/go/gotenv",
		RepoURL:      "https://github.com/subosito/gotenv",
	}

	str := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<title>subosito.com/go/gotenv</title>
	<meta name="go-import" content="subosito.com/go/gotenv git https://github.com/subosito/gotenv">
	<meta name="go-source" content="subosito.com/go/gotenv https://github.com/subosito/gotenv https://github.com/subosito/gotenv/tree/master{/dir} https://github.com/subosito/gotenv/blob/master{/dir}/{file}#L{line}">
	<meta http-equiv="refresh" content="0; url=https://godoc.org/subosito.com/go/gotenv">
</head>
<body>
	Nothing to see here. Please <a href="https://godoc.org/subosito.com/go/gotenv">move along</a>.
</body>
</html>
`

	out := &strings.Builder{}
	tiffany.Render(out, opt)

	assert.Equal(t, out.String(), str)
}
