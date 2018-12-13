package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/deanishe/awgo"
	"github.com/xruins/alfred3-workflow-godoc/godoc"
)

var (
	helpURL    = "https://github.com/xruins/alfred3-workflow-godoc"
	maxResults = 200
	wf         *aw.Workflow
	icon       *aw.Icon
)

func init() {
	wf = aw.New(aw.HelpURL(helpURL), aw.MaxResults(maxResults))

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	icon = &aw.Icon{
		Value: dir + "icon.png",
		Type:  aw.IconTypeFileType,
	}
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
			Arg(r.Path).
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
