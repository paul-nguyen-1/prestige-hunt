package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type job struct {
	title    string
	category string
	url      string
}

var hardcodedKeywords = [...]string{"software engineer intern", "software developer intern", "intern"}

func scrapeJobs() {
	c := colly.NewCollector(
		colly.AllowedDomains("jobs.lever.co"),
	)

	c.OnHTML("a[class=posting-title]", func(h *colly.HTMLElement) {
		leverJob := job{
			title:    h.ChildText("h5[data-qa=posting-name]"),
			category: h.ChildText("div.posting-categories"),
			url:      h.Attr("href"),
		}
		if containsKeyword(leverJob.title, hardcodedKeywords[:]) {
			fmt.Println(leverJob.title + " | " + leverJob.category + " | " + leverJob.url)
		}
	})

	c.Visit("https://jobs.lever.co/boringcompany")
}

func containsKeyword(title string, keywords []string) bool {
	for _, key := range keywords {
		if strings.Contains(strings.ToLower(title), key) {
			return true
		}
	}
	return false
}
