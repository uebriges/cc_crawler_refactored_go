package crawler

import (
	"errors"
	"net/url"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func NormalizeLink(baseURL, link string) (string, error) {
	link = strings.Trim(link, " 	")
	switch {
	case strings.HasPrefix(link, "/"):
		return baseURL + link, nil
	case strings.HasPrefix(link, "http://"),
		strings.HasPrefix(link, "https://"):
		return link, nil
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	p := u.Path
	if strings.HasSuffix(p, ".html") {
		parts := strings.Split(p, "/")
		p = strings.Join(parts[0:len(parts)-1], "/")
	}
	pathParts := strings.Split(strings.TrimPrefix(p, "/"), "/") // /foo/bar
	linkParts := strings.Split(link, "/")                       // ../../bla/blubb

	i := 0
	for _, linkPart := range linkParts {
		if linkPart == "" {
			continue
		}
		if linkPart != ".." {
			break
		}
		i++
	}

	pathEnd := len(pathParts) - i
	if pathEnd < 1 {
		spew.Dump(baseURL, link)
		return "", errors.New("invalid path end")
	}
	path := append(pathParts[0:pathEnd], linkParts[i:]...)

	u.Path = strings.Join(path, "/")
	return u.String(), nil

}
