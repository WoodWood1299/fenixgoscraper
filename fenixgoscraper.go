package fenixgoscraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Announcement struct {
	Link string
}

func ScrapeOut() string {
	return "I'M SCRAPING\n"
}

func Scrape() {
	var announcements []Announcement

	c := colly.NewCollector(
		colly.AllowedDomains("fenix.tecnico.ulisboa.pt/disciplinas"))

	c.OnHTML("h5[style=margin-top: 0; font-weight:400]", func(h *colly.HTMLElement) {
		announcement := Announcement{}

		announcement.Link = h.ChildAttr("a", "href")

		announcements = append(announcements, announcement)
	})

	c.OnScraped(func(r *colly.Response) {
		for i := 0; i < len(announcements); i++ {
			fmt.Print(announcements[i].Link)
		}
	})

	c.Visit("fenix.tecnico.ulisboa.pt/disciplinas/OC112/2025-2026/1-semestre/laboratorios")
}
