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

func CreateAnnouncement(Item *gofeed.Item) Announcement {
	var a Announcement
	a.Link = Item.Link
	a.Message = Item.Title
	return a
}

func StringAnnouncement(announcement Announcement) string {
	return fmt.Sprintf("%s\n%s\n\n", html.UnescapeString(announcement.Message), announcement.Link)
}

func Scrape(disciplina_links map[string]string, announcement_count int) (map[string][]Announcement, error) {

	if len(disciplina_links) == 0 {
		return nil, errors.New("link array cannot be empty")
	}

	announcements := make(map[string][]Announcement, len(disciplina_links))

	for disciplina := range disciplina_links {
		announcements[disciplina] = make([]Announcement, announcement_count)
	}

	fp := gofeed.NewParser()

	for disciplina, link := range disciplina_links {
		feed, err := fp.ParseURL(link)

		if err != nil {
			return nil, errors.New("error parsing RSS")
		}

		items := feed.Items
		count := min(announcement_count, len(items))
		for j := 0; j < count; j++ {
			announcements[disciplina][j] = CreateAnnouncement(items[j])
		}
	}

	return announcements, nil
}
