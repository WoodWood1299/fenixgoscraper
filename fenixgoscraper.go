package fenixgoscraper

import (
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

func Scrape() {
	fmt.Print("Running\n")
	links := [4]string{
		"https://fenix.tecnico.ulisboa.pt/disciplinas/OC112/2025-2026/1-semestre/rss/announcement",
		"https://fenix.tecnico.ulisboa.pt/disciplinas/Apre2222/2025-2026/1-semestre/rss/announcement",
		"https://fenix.tecnico.ulisboa.pt/disciplinas/OC112/2025-2026/1-semestre/rss/announcement",
		"https://fenix.tecnico.ulisboa.pt/disciplinas/RC112/2025-2026/1-semestre/rss/announcement",
	}

	fp := gofeed.NewParser()

	for _, l := range links {
		feed, _ := fp.ParseURL(l)
		fmt.Printf("%s\n", feed.Title)

		items := feed.Items
		for _, a := range items {
			fmt.Printf("- %s\n\t%s\n\n", html.UnescapeString(a.Title), a.Link)
		}
	}
	fmt.Print("Done\n")
}
