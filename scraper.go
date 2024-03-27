package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func scrapeJobs() {
	c := colly.NewCollector(
		colly.AllowedDomains("jobs.lever.co"),
	)

	c.OnHTML("a[class=posting-title]", func(h *colly.HTMLElement) {
		job := h.ChildText("h5[data-qa=posting-name]")
		categories := h.ChildText("div.posting-categories")
		link := h.Attr("href")

		fmt.Println(job + " | " + categories + " | " + link)
	})

	c.Visit("https://jobs.lever.co/boringcompany")
}
