package tiffany_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subosito/tiffany"
)

func TestRender(t *testing.T) {
	data := []struct {
		name     string
		option   tiffany.Option
		expected string
	}{
		{
			name: "github",
			option: tiffany.Option{
				CanonicalURL: "subosito.com/go/gotenv",
				RepoURL:      "https://github.com/subosito/gotenv",
			},
			expected: `
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
`,
		},
		{
			name: "bitbucket",
			option: tiffany.Option{
				CanonicalURL: "subosito.com/go/gotenv",
				RepoURL:      "https://bitbucket.org/subosito/gotenv",
			},
			expected: `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<title>subosito.com/go/gotenv</title>
	<meta name="go-import" content="subosito.com/go/gotenv git https://bitbucket.org/subosito/gotenv">
	<meta name="go-source" content="subosito.com/go/gotenv https://bitbucket.org/subosito/gotenv https://bitbucket.org/subosito/gotenv/src/default{/dir} https://bitbucket.org/subosito/gotenv/src/default{/dir}/{file}#{file}-{line}">
	<meta http-equiv="refresh" content="0; url=https://godoc.org/subosito.com/go/gotenv">
</head>
<body>
	Nothing to see here. Please <a href="https://godoc.org/subosito.com/go/gotenv">move along</a>.
</body>
</html>
`,
		},
		{
			name: "gitlab + custom godoc",
			option: tiffany.Option{
				CanonicalURL: "subosito.com/go/gotenv",
				RepoURL:      "https://gitlab.com/subosito/gotenv",
				VCS:          "git",
				GodocURL:     "https://doc.example.com",
			},
			expected: `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<title>subosito.com/go/gotenv</title>
	<meta name="go-import" content="subosito.com/go/gotenv git https://gitlab.com/subosito/gotenv">
	
	<meta http-equiv="refresh" content="0; url=https://doc.example.com/subosito.com/go/gotenv">
</head>
<body>
	Nothing to see here. Please <a href="https://doc.example.com/subosito.com/go/gotenv">move along</a>.
</body>
</html>
`,
		},
	}

	for i := range data {
		t.Run(data[i].name, func(t *testing.T) {
			out := &strings.Builder{}
			tiffany.Render(out, data[i].option)
			assert.Equal(t, data[i].expected, out.String())
		})
	}
}
