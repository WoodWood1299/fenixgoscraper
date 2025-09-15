package fenixgoscraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Announcement struct {
	Link, Message string
}

func ScrapeOut() {
	fmt.Print("TEST\n")
}

func onHTML(announcements *[]Announcement, h *colly.HTMLElement) {
	announcement := Announcement{}
	announcement.Link = h.ChildAttr("a", "href")
	announcement.Message = h.Text
	*announcements = append(*announcements, announcement)
}

func onScrape(announcements *[]Announcement) {
	for _, a := range *announcements {
		fmt.Printf("- %s\n\t%s\n\n", a.Message, a.Link)
	}
}

func scraperSetup(announcements *[]Announcement) *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("fenix.tecnico.ulisboa.pt"),
	)

	c.OnHTML("h5[style]", func(h *colly.HTMLElement) {
		onHTML(announcements, h)
	})

	c.OnScraped(func(r *colly.Response) {
		onScrape(announcements)
	})

	return c
}

func Scrape() {
	fmt.Print("Running\n")
	links := [4]string{
		"https://fenix.tecnico.ulisboa.pt/disciplinas/Apre2222/2025-2026/1-semestre/anuncios",
		"https://fenix.tecnico.ulisboa.pt/disciplinas/OC112/2025-2026/1-semestre/anuncios",
		"https://fenix.tecnico.ulisboa.pt/disciplinas/Mod112/2025-2026/1-semestre/anuncios",
		"https://fenix.tecnico.ulisboa.pt/disciplinas/RC112/2025-2026/1-semestre/anuncios",
	}

	var announcements []Announcement
	c := scraperSetup(&announcements)

	for _, link := range links {
		fmt.Println("XXX VISITING XXX")
		c.Visit(link)
		announcements = nil
	}
	fmt.Print("Done\n")
}
