package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	// Scrape by individual company
	// go scrapeLeverCompany("netflix")

	// Scrape companies concurrently
	var waitGroup sync.WaitGroup
	jobChannel := make(chan job)

	go processJobs(jobChannel)
	for _, company := range hardcodedLeverCompanies {
		waitGroup.Add(1)
		go scrapeLeverCompanyList(company, &waitGroup, jobChannel)
	}

	waitGroup.Wait()
	close(jobChannel)

	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen to server", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
