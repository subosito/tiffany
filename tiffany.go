package tiffany // import "subosito.com/go/tiffany"

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
{{- if not .RedirectDisabled}}
	<title>{{.CanonicalURL}}</title>
{{- end}}
	<meta name="go-import" content="{{.CanonicalURL}} {{.VCS}} {{.RepoURL}}">
{{if .SourceURL}}	<meta name="go-source" content="{{.CanonicalURL}} {{.RepoURL}} {{.SourceURL}}">{{- end}}
{{- if not .RedirectDisabled}}
	<meta http-equiv="refresh" content="0; url={{if .GodocDisabled}}{{.RedirectURL}}{{else}}{{.GodocURL}}/{{.CanonicalURL}}{{end}}">
{{end -}}
</head>
{{- if not .RedirectDisabled}}
<body>
	Nothing to see here. Please <a href="{{if .GodocDisabled}}{{.RedirectURL}}{{else}}{{.GodocURL}}/{{.CanonicalURL}}{{end}}">move along</a>.
</body>
{{- end}}
</html>
`))

const (
	godocURL     = "https://godoc.org"
	githubURL    = "https://github.com"
	bitbucketURL = "https://bitbucket.org"
)

// Option is configuration for vanity URL
type Option struct {
	CanonicalURL     string
	RepoURL          string
	VCS              string
	SourceLayout     string
	SourceURL        string
	GodocURL         string
	GodocDisabled    bool
	RedirectURL      string
	RedirectDisabled bool
}

func (opt Option) vcs() string {
	switch {
	case strings.HasPrefix(opt.RepoURL, githubURL):
		return "git"
	case strings.HasPrefix(opt.RepoURL, bitbucketURL) && opt.VCS == "":
		return "git"
	default:
		return opt.VCS
	}
}

func (opt Option) sourceLayout() string {
	if opt.SourceLayout != "" {
		return opt.SourceLayout
	}

	switch {
	case strings.HasPrefix(opt.RepoURL, githubURL):
		return "github"
	case strings.HasPrefix(opt.RepoURL, bitbucketURL):
		return "bitbucket"
	default:
		return ""
	}
}

func (opt Option) sourceURL() string {
	layout := opt.sourceLayout()

	switch layout {
	case "":
		return ""
	case "github":
		return fmt.Sprintf("%v/tree/master{/dir} %v/blob/master{/dir}/{file}#L{line}", opt.RepoURL, opt.RepoURL)
	case "bitbucket":
		return fmt.Sprintf("%v/src/default{/dir} %v/src/default{/dir}/{file}#{file}-{line}", opt.RepoURL, opt.RepoURL)
	case "gogs":
		return fmt.Sprintf("%v/src/master{/dir} %v/src/master{/dir}/{file}#L{line}", opt.RepoURL, opt.RepoURL)
	default:
		return fmt.Sprintf(layout, opt.RepoURL, opt.RepoURL)
	}
}

func (opt Option) godocURL() string {
	if opt.GodocURL == "" {
		return godocURL
	}

	return opt.GodocURL
}

func (opt Option) redirectURL() string {
	if opt.RedirectURL != "" {
		return opt.RedirectURL
	}

	return opt.RepoURL
}

// Render renders the vanity URL information based on supplied option.
func Render(w io.Writer, option Option) error {
	return vanityTmpl.Execute(w, Option{
		CanonicalURL:     option.CanonicalURL,
		RepoURL:          option.RepoURL,
		VCS:              option.vcs(),
		SourceURL:        option.sourceURL(),
		GodocURL:         option.godocURL(),
		GodocDisabled:    option.GodocDisabled,
		RedirectURL:      option.redirectURL(),
		RedirectDisabled: option.RedirectDisabled,
	})
}
