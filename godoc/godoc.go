package godoc

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/antchfx/xquery/html"
	"golang.org/x/net/html"
)

const goDocURL = "https://godoc.org/"

// Result represent searching results of godoc.org
type Result struct {
	Path     string
	Synopsis string
}

// request requests godoc.org with given query, without following redirection
func requestWithoutRedirect(query string) (*http.Response, error) {
	url := fmt.Sprintf("%s?q=%s", goDocURL, url.QueryEscape(query))
	client := http.DefaultClient
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return client.Get(url)
}

// parseDoc parses given HTML, then returns its package description
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

// parseHTML parses given HTML, then returns a slice of *Result
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

		sn := htmlquery.Find(p, "/td[@class='synopsis']")
		var synopsysNode *html.Node
		if len(sn) == 1 {
			synopsysNode = sn[0]
		}
		// trim zero-width space
		trimed := strings.Replace(htmlquery.InnerText(pathNode), "\u200b", "", -1)
		res := &Result{
			Path:     goDocURL + trimed,
			Synopsis: htmlquery.InnerText(synopsysNode),
		}
		//runtime.Breakpoint()
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

	// godoc.org returns HTTP status code 302 for exact match such as "https://godoc.org/?q=net/http".
	// then, parse godoc to get synopsis.
	if res.StatusCode == http.StatusFound {
		synopsis, err := parseDoc(res.Body)
		if err != nil {
			return nil, err
		}

		result := &Result{
			Path:     goDocURL + res.Header.Get("Location"),
			Synopsis: synopsis,
		}
		return []*Result{result}, nil
	} else {
		return parseSearchResult(res.Body)
	}
}
