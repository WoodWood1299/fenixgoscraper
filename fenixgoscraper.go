package fenixgoscraper

import (
	"errors"
	"fmt"

	"github.com/mmcdole/gofeed"
)

type Announcement struct {
	Link, Message string
}

func ScrapeOut() {
	fmt.Print("TEST\n")
}

func createAnnouncement(Item *gofeed.Item) Announcement {
	var a Announcement
	a.Link = Item.Link
	a.Message = Item.Title
	return a
}

func Scrape(links []string, announcement_count int) ([][]Announcement, error) {

	if len(links) == 0 {
		return nil, errors.New("link array cannot be empty")
	}

	announcements := make([][]Announcement, len(links))
	for i := range announcements {
		announcements[i] = make([]Announcement, announcement_count)
	}

	fp := gofeed.NewParser()

	for i, l := range links {
		feed, err := fp.ParseURL(l)

		if err != nil {
			return nil, errors.New("error parsing RSS")
		}

		//out += fmt.Sprintf("%s\n", feed.Title)

		items := feed.Items
		for j := 0; j < announcement_count; j++ {
			announcements[i][j] = createAnnouncement(items[j])
			//out += fmt.Sprintf("- %s\n\t%s\n\n", html.UnescapeString(items[i].Title), items[i].Link)
		}
	}

	return announcements, nil
}
