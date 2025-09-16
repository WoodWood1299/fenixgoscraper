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

func Scrape(links []string, announcement_count int) (string, error) {

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
		for i := 0; i < announcement_count; i++ {
			out += fmt.Sprintf("- %s\n\t%s\n\n", html.UnescapeString(items[i].Title), items[i].Link)
		}
	}

	return out, nil
}
