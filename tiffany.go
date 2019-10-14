package tiffany // import "subosito.com/go/tiffany"

import (
	"fmt"
	"html/template"
	"io"
	"net/url"
	"strings"
)

var vanityTmpl = template.Must(template.New("vanity").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<title>{{.CanonicalURL}}</title>
	<meta name="go-import" content="{{.CanonicalURL}} {{.VCS}} {{.RepoURL}}">
{{if .SourceURL}}	<meta name="go-source" content="{{.CanonicalURL}} {{.RepoURL}} {{.SourceURL}}">{{- end}}
	<meta http-equiv="refresh" content="0; url={{if .GodocDisabled}}{{.RepoURL}}{{else}}{{.GodocURL}}/{{.CanonicalURL}}{{end}}">
</head>
<body>
	Nothing to see here. Please <a href="{{if .GodocDisabled}}{{.RepoURL}}{{else}}{{.GodocURL}}/{{.CanonicalURL}}{{end}}">move along</a>.
</body>
</html>
`))

const (
	godocURL     = "https://godoc.org"
	githubURL    = "https://github.com"
	bitbucketURL = "https://bitbucket.org"
)

// Option is configuration for vanity URL
type Option struct {
	CanonicalURL  string
	RepoURL       string
	VCS           string
	SourceLayout  string
	SourceURL     string
	GodocURL      string
	GodocDisabled bool
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

func (opt Option) repoURL() (string, error) {
	u, err := url.Parse(opt.RepoURL)
	if err != nil {
		return "", err
	}

	var b strings.Builder

	if strings.HasPrefix(u.Scheme, "http") {
		b.WriteString(u.Scheme)
	} else {
		b.WriteString("https")
	}

	b.WriteString("://")
	b.WriteString(u.Host)
	b.WriteString(u.Path)

	return b.String(), nil
}

// Render renders the vanity URL information based on supplied option.
func Render(w io.Writer, option Option) error {
	repoURL, err := option.repoURL()
	if err != nil {
		return err
	}

	return vanityTmpl.Execute(w, Option{
		CanonicalURL:  option.CanonicalURL,
		RepoURL:       repoURL,
		VCS:           option.vcs(),
		SourceURL:     option.sourceURL(),
		GodocURL:      option.godocURL(),
		GodocDisabled: option.GodocDisabled,
	})
}
