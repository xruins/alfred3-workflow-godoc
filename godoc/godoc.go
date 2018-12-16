package godoc

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/antchfx/xquery/html"
	"golang.org/x/net/html"
)

const goDocURL = "https://godoc.org"

// Result represents the searching results of godoc.org
type Result struct {
	Path     string
	URL      *url.URL
	Synopsis string
}

func makeResult(url *url.URL, synopsis string) *Result {
	return &Result{
		Path:     url.Path,
		URL:      url,
		Synopsis: synopsis,
	}
}

// request requests godoc.org with given query, without following redirection
func requestWithoutRedirect(query string) (*http.Response, error) {
	u, err := url.Parse(goDocURL)
	if err != nil {
		return nil, err
	}
	q := fmt.Sprintf("?q=%s", url.QueryEscape(query))
	u.RawQuery = q
	client := http.DefaultClient
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return client.Get(u.String())
}

// parseDoc parses given HTML, then returns its package description
// an argument should be document of godoc.org (e.g. https://godoc.org/net/)
func parseDoc(r io.Reader) (string, error) {
	doc, err := htmlquery.Parse(r)
	if err != nil {
		return "", err
	}
	// get package description as synopsys
	n := htmlquery.Find(doc, "/html/body/div[1]/p[2]")
	if len(n) == 1 {
		return htmlquery.InnerText(n[0]), nil
	}
	return "", nil
}

// parseSearchResult parses given HTML, then returns a slice of *Result
// an argument should be search result of godoc.org (e.g. https://godoc.org/?q=foo)
func parseSearchResult(r io.Reader) ([]*Result, error) {
	doc, err := htmlquery.Parse(r)
	if err != nil {
		return nil, err
	}

	var ret []*Result
	parentNodes := htmlquery.Find(doc, "/html/body/div/table/tbody/tr")
	for _, p := range parentNodes {
		pn := htmlquery.Find(p, "/td/a")
		// skip if tr tag does not include link
		if len(pn) != 1 {
			continue
		}
		pathNode := pn[0]
		relPath := htmlquery.SelectAttr(pathNode, "href")

		sn := htmlquery.Find(p, "/td[@class='synopsis']")
		var synopsysNode *html.Node
		if len(sn) == 1 {
			synopsysNode = sn[0]
		}
		u, err := url.Parse(goDocURL)
		if err != nil {
			return nil, err
		}
		u.Path = path.Join(u.Path, relPath)
		synopsis := htmlquery.InnerText(synopsysNode)
		res := makeResult(u, synopsis)
		ret = append(ret, res)
	}

	return ret, nil
}

// Search searches godoc.org with given query and returns found URLs
func Search(query string) ([]*Result, error) {
	res, err := requestWithoutRedirect(query)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// if the query matches exact package name, parse document to get synopsis
	if res.StatusCode == http.StatusFound {
		u, err := url.Parse(goDocURL + res.Header.Get("Location"))
		if err != nil {
			return nil, err
		}
		synopsis, err := parseDoc(res.Body)
		if err != nil {
			return nil, err
		}

		result := makeResult(u, synopsis)
		return []*Result{result}, nil
	}
	return parseSearchResult(res.Body)
}
