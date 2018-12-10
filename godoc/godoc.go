package godoc

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/antchfx/xquery/html"
)

const goDocUrl = "https://godoc.org"

type Result struct {
	Path     string
	Synopsis string
}

// Search searches godoc.org with given query and returns found URLs
func Search(query string) ([]*Result, error) {

	url := fmt.Sprintf("%s?q=%s", goDocUrl, url.QueryEscape(query))
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := htmlquery.Parse(res.Body)
	if err != nil {
		return nil, err
	}

	var ret []*Result
	// get synopsis
	synopses := htmlquery.Find(doc, "/html/body/div/table/tbody/tr/td.synopsys")

	for i, n := range htmlquery.Find(doc, "/html/body/div/table/tbody/tr/td/a") {
		r := &Result{
			Path:     goDocUrl + htmlquery.InnerText(n),
			Synopsis: htmlquery.InnerText(synopses[i]),
		}
		ret = append(ret, r)
	}

	return ret, nil
}
