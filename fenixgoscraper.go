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

	for i, link := range links {
		feed, err := fp.ParseURL(link)

		if err != nil {
			return nil, errors.New("error parsing RSS")
		}

		items := feed.Items
		count := min(announcement_count, len(items))
		for j := 0; j < count; j++ {
			announcements[i][j] = createAnnouncement(items[j])
		}
	}

	return announcements, nil
}
