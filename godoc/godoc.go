package godoc

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/antchfx/xquery/html"
)

const goDocUrl = "https://godoc.org"

// Result represent searching results of godoc.org
type Result struct {
	Path     string
	Synopsis string
}

// request requests godoc.org with given query
func request(query string) (*http.Response, error) {
	url := fmt.Sprintf("%s?q=%s", goDocUrl, url.QueryEscape(query))
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
	parentNodes := htmlquery.Find(doc, "/html/body/div/table/tbody/tr/td")

	for _, p := range parentNodes {
		pathNode := htmlquery.Find(p, "/td/a")[0]
		synopsysNode := htmlquery.Find(p, "/td[@class='synopsis']")[0]
		r := &Result{
			Path:     goDocUrl + htmlquery.InnerText(pathNode),
			Synopsis: htmlquery.InnerText(synopsysNode),
		}
		ret = append(ret, r)
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
