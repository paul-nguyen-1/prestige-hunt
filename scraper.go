package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

type job struct {
	title    string
	category string
	url      string
}

var hardcodedKeywords = [...]string{"software engineer intern", "software developer intern", "intern"}
var hardcodedLeverCompanies = [...]string{"netflix", "boringcompany", "palantir"}

// func scrapeLeverCompany(company string) {
// 	c := colly.NewCollector(
// 		colly.AllowedDomains("jobs.lever.co"),
// 	)

// 	c.OnHTML("a[class=posting-title]", func(h *colly.HTMLElement) {
// 		leverJob := job{
// 			title:    h.ChildText("h5[data-qa=posting-name]"),
// 			category: h.ChildText("div.posting-categories"),
// 			url:      h.Attr("href"),
// 		}
// 		if containsKeyword(leverJob.title, hardcodedKeywords[:]) {
// 			fmt.Println(leverJob.title + " | " + leverJob.category + " | " + leverJob.url)
// 		}
// 	})

// 	url := fmt.Sprintf("https://jobs.lever.co/%s", company)
// 	c.Visit(url)
// }

func scrapeLeverCompanyList(company string, waitGroup *sync.WaitGroup, jobChannel chan<- job) {
	defer waitGroup.Done()
	c := colly.NewCollector(
		colly.AllowedDomains("jobs.lever.co"),
	)

	c.OnHTML("a[class=posting-title]", func(h *colly.HTMLElement) {
		leverJob := job{
			title:    h.ChildText("h5[data-qa=posting-name]"),
			category: h.ChildText("div.posting-categories"),
			url:      h.Attr("href"),
		}
		// Send job to job channel for processing
		jobChannel <- leverJob
	})

	url := fmt.Sprintf("https://jobs.lever.co/%s", company)
	c.Visit(url)
}

func processJobs(jobChannel <-chan job) {
	for job := range jobChannel {
		if containsKeyword(job.title, hardcodedKeywords[:]) {
			fmt.Printf("%s | %s | %s\n", job.title, job.category, job.url)
		}
	}
}

func containsKeyword(title string, keywords []string) bool {
	for _, key := range keywords {
		if strings.Contains(strings.ToLower(title), key) {
			return true
		}
	}
	return false
}
