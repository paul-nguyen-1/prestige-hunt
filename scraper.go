package main

import (
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("https://jobs.lever.co/"),
	)

	c.Visit("https://jobs.lever.co")
}
