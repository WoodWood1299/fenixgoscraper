package fenixgoscraper

import (
	"errors"
	"fmt"
	"html"

	"github.com/mmcdole/gofeed"
)

type Announcement struct {
	Link, Message string
}

func ScrapeOut() {
	fmt.Print("TEST\n")
}

func Scrape(links []string) (string, error) {

	if len(links) == 0 {
		return "", errors.New("link array cannot be empty")
	}

	var out string
	fp := gofeed.NewParser()

	for _, l := range links {
		feed, err := fp.ParseURL(l)

		if err != nil {
			return "", errors.New("error parsing RSS")
		}

		out += fmt.Sprintf("%s\n", feed.Title)

		items := feed.Items
		//for _, a := range items {
		//	out += fmt.Sprintf("- %s\n\t%s\n\n", html.UnescapeString(a.Title), a.Link)
		//}

		out += fmt.Sprintf("- %s\n\t%s\n\n", html.UnescapeString(items[0].Title), items[0].Link)
	}

	return out, nil
}
