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

// request requests godoc.org with given query
func request(query string) (*http.Response, error) {
	url := fmt.Sprintf("%s?q=%s", goDocURL, url.QueryEscape(query))
	return http.Get(url)
}

// parseHTML parses given HTML, then returns a slice of *Result
func parseHTML(r io.Reader) ([]*Result, error) {
	doc, err := htmlquery.Parse(r)
	if err != nil {
		return nil, err
	}

	var ret []*Result
	// get synopsis
	parentNodes := htmlquery.Find(doc, "/html/body/div/table/tbody/tr")
	for _, p := range parentNodes {
		pn := htmlquery.Find(p, "/td/a")
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
	res, err := request(query)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return parseHTML(res.Body)
}
