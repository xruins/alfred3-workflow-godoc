package main

import (
	"github.com/deanishe/awgo"
	"github.com/xruins/alfred3-workflow-godoc/godoc"
)

var (
	helpURL    = "https://github.com/xruins/alfred3-workflow-godoc"
	maxResults = 200
	wf         *aw.Workflow

	// Icon for bookmark filetype
	icon = &aw.Icon{
		Value: "mobi.ruins.alfred3-workflow-godoc",
		Type:  aw.IconTypeFileType,
	}
)

func init() {
	wf = aw.New(aw.HelpURL(helpURL), aw.MaxResults(maxResults))
}

func run() {
	var query string

	// Use wf.Args() to enable Magic Actions
	if args := wf.Args(); len(args) > 0 {
		query = args[0]
	}
	results, err := godoc.Search(query)
	if err != nil {
		wf.FatalError(err)
	}

	for _, r := range results {

		wf.NewItem(r.Path).
			Subtitle(r.Synopsis).
			Arg(r.Url.String()).
			UID(r.Path).
			Icon(icon).
			Valid(true)
	}

	wf.WarnEmpty("No matching results", "Try a different query?")
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
