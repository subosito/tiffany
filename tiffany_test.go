package tiffany_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"subosito.com/go/tiffany"
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
			name: "github + custom godoc",
			option: tiffany.Option{
				CanonicalURL: "subosito.com/go/gotenv",
				RepoURL:      "https://github.com/subosito/gotenv",
				GodocURL:     "https://doc.example.com",
			},
			expected: `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<title>subosito.com/go/gotenv</title>
	<meta name="go-import" content="subosito.com/go/gotenv git https://github.com/subosito/gotenv">
	<meta name="go-source" content="subosito.com/go/gotenv https://github.com/subosito/gotenv https://github.com/subosito/gotenv/tree/master{/dir} https://github.com/subosito/gotenv/blob/master{/dir}/{file}#L{line}">
	<meta http-equiv="refresh" content="0; url=https://doc.example.com/subosito.com/go/gotenv">
</head>
<body>
	Nothing to see here. Please <a href="https://doc.example.com/subosito.com/go/gotenv">move along</a>.
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
			name: "gogs",
			option: tiffany.Option{
				CanonicalURL: "subosito.com/go/gotenv",
				RepoURL:      "https://git.subosito.com/subosito/gotenv",
				SourceLayout: "gogs",
				VCS:          "git",
			},
			expected: `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<title>subosito.com/go/gotenv</title>
	<meta name="go-import" content="subosito.com/go/gotenv git https://git.subosito.com/subosito/gotenv">
	<meta name="go-source" content="subosito.com/go/gotenv https://git.subosito.com/subosito/gotenv https://git.subosito.com/subosito/gotenv/src/master{/dir} https://git.subosito.com/subosito/gotenv/src/master{/dir}/{file}#L{line}">
	<meta http-equiv="refresh" content="0; url=https://godoc.org/subosito.com/go/gotenv">
</head>
<body>
	Nothing to see here. Please <a href="https://godoc.org/subosito.com/go/gotenv">move along</a>.
</body>
</html>
`,
		},
		{
			name: "gitlab",
			option: tiffany.Option{
				CanonicalURL: "subosito.com/go/gotenv",
				RepoURL:      "https://gitlab.com/subosito/gotenv",
				VCS:          "git",
			},
			expected: `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<title>subosito.com/go/gotenv</title>
	<meta name="go-import" content="subosito.com/go/gotenv git https://gitlab.com/subosito/gotenv">

	<meta http-equiv="refresh" content="0; url=https://godoc.org/subosito.com/go/gotenv">
</head>
<body>
	Nothing to see here. Please <a href="https://godoc.org/subosito.com/go/gotenv">move along</a>.
</body>
</html>
`,
		},
		{
			name: "gitlab + without godoc",
			option: tiffany.Option{
				CanonicalURL:  "subosito.com/go/gotenv",
				RepoURL:       "https://gitlab.com/subosito/gotenv",
				GodocDisabled: true,
				VCS:           "git",
			},
			expected: `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<title>subosito.com/go/gotenv</title>
	<meta name="go-import" content="subosito.com/go/gotenv git https://gitlab.com/subosito/gotenv">

	<meta http-equiv="refresh" content="0; url=https://gitlab.com/subosito/gotenv">
</head>
<body>
	Nothing to see here. Please <a href="https://gitlab.com/subosito/gotenv">move along</a>.
</body>
</html>
`,
		},
		{
			name: "gitlab ssh + custom redirection",
			option: tiffany.Option{
				CanonicalURL:  "subosito.com/go/gotenv",
				RepoURL:       "ssh://git@git.gitlab.com/subosito/gotenv",
				RedirectURL:   "https://gitlab.com/subosito/gotenv",
				GodocDisabled: true,
				VCS:           "git",
			},
			expected: `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<title>subosito.com/go/gotenv</title>
	<meta name="go-import" content="subosito.com/go/gotenv git ssh://git@git.gitlab.com/subosito/gotenv">

	<meta http-equiv="refresh" content="0; url=https://gitlab.com/subosito/gotenv">
</head>
<body>
	Nothing to see here. Please <a href="https://gitlab.com/subosito/gotenv">move along</a>.
</body>
</html>
`,
		},
		{
			name: "gitlab ssh + disable redirection",
			option: tiffany.Option{
				CanonicalURL:     "subosito.com/go/gotenv",
				RepoURL:          "ssh://git@git.gitlab.com/subosito/gotenv",
				GodocDisabled:    true,
				RedirectDisabled: true,
				VCS:              "git",
			},
			expected: `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<title>subosito.com/go/gotenv</title>
	<meta name="go-import" content="subosito.com/go/gotenv git ssh://git@git.gitlab.com/subosito/gotenv">
</head>
</html>
`,
		},
		{
			name: "custom",
			option: tiffany.Option{
				CanonicalURL: "subosito.com/go/gotenv",
				RepoURL:      "https://git.subosito.com/subosito/gotenv",
				SourceLayout: "%v/dirs/master{/dir} %v/files/master{/dir}/{file}#L{line}",
				VCS:          "git",
			},
			expected: `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<title>subosito.com/go/gotenv</title>
	<meta name="go-import" content="subosito.com/go/gotenv git https://git.subosito.com/subosito/gotenv">
	<meta name="go-source" content="subosito.com/go/gotenv https://git.subosito.com/subosito/gotenv https://git.subosito.com/subosito/gotenv/dirs/master{/dir} https://git.subosito.com/subosito/gotenv/files/master{/dir}/{file}#L{line}">
	<meta http-equiv="refresh" content="0; url=https://godoc.org/subosito.com/go/gotenv">
</head>
<body>
	Nothing to see here. Please <a href="https://godoc.org/subosito.com/go/gotenv">move along</a>.
</body>
</html>
`,
		},
	}

	for _, val := range data {
		t.Run(val.name, func(t *testing.T) {
			out := &strings.Builder{}
			tiffany.Render(out, val.option)
			assert.Equal(t, val.expected, out.String())
		})
	}
}
