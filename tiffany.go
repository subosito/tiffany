package tiffany

import (
	"fmt"
	"html/template"
	"io"
	"strings"
)

var vanityTmpl = template.Must(template.New("vanity").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<title>{{.CanonicalURL}}</title>
	<meta name="go-import" content="{{.CanonicalURL}} {{.VCS}} {{.RepoURL}}">
	{{if .SourceURL -}}
	<meta name="go-source" content="{{.CanonicalURL}} {{.RepoURL}} {{.SourceURL}}">
	{{- end}}
	<meta http-equiv="refresh" content="0; url={{.GodocURL}}/{{.CanonicalURL}}">
</head>
<body>
	Nothing to see here. Please <a href="{{.GodocURL}}/{{.CanonicalURL}}">move along</a>.
</body>
</html>
`))

const (
	godocURL     = "https://godoc.org"
	githubURL    = "https://github.com"
	bitbucketURL = "https://bitbucket.org"
)

type Option struct {
	CanonicalURL string
	RepoURL      string
	VCS          string
	SourceURL    string
	GodocURL     string
}

func (opt Option) vcs() string {
	switch {
	case strings.HasPrefix(opt.RepoURL, githubURL), strings.HasPrefix(opt.RepoURL, bitbucketURL):
		return "git"
	default:
		return opt.VCS
	}
}

func (opt Option) sourceURL() string {
	switch {
	case strings.HasPrefix(opt.RepoURL, githubURL):
		return fmt.Sprintf("%v/tree/master{/dir} %v/blob/master{/dir}/{file}#L{line}", opt.RepoURL, opt.RepoURL)
	case strings.HasPrefix(opt.RepoURL, bitbucketURL):
		return fmt.Sprintf("%v/src/default{/dir} %v/src/default{/dir}/{file}#{file}-{line}", opt.RepoURL, opt.RepoURL)
	default:
		return ""
	}
}

func (opt Option) godocURL() string {
	if opt.GodocURL == "" {
		return godocURL
	}

	return opt.GodocURL
}

func Render(w io.Writer, option Option) error {
	return vanityTmpl.Execute(w, Option{
		CanonicalURL: option.CanonicalURL,
		RepoURL:      option.RepoURL,
		VCS:          option.vcs(),
		SourceURL:    option.sourceURL(),
		GodocURL:     option.godocURL(),
	})
}
