package main

import (
		"fmt"
		"encoding/json"
		"io/ioutil"
		"os"
		"log"
		"github.com/gocolly/colly"
		"strings"
)

// Repo struct represents what info will be scraped from github trending page
type Repo struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
}

// Function that creates and writes info to the json file
func createJsonFile(filename string, repos []Repo) {
	jsonData, err := json.MarshalIndent(repos, "", "    ")

	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(filename, jsonData, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

// Function that scrpaes name and description from github trending page
func scrapeHandler() {
	c := colly.NewCollector()
	repos := make([]Repo, 0)

	// On every a element which has title attribute call callback
	c.OnHTML("body > div.application-main > main > div.explore-pjax-container.container-lg.p-responsive.pt-6 > div > div:nth-child(2) > article:nth-child(n)", func(e *colly.HTMLElement) {
		// Clean up data before assigning to struct
		repo := Repo{}

		info := e.ChildText("body > div.application-main > main > div.explore-pjax-container.container-lg.p-responsive.pt-6 > div > div:nth-child(2) > article:nth-child(n) > h1")
		name := strings.Replace(info, "\n\n     ", "", -1)
		repo.Name = name

		repo.Description = e.ChildText("body > div.application-main > main > div.explore-pjax-container.container-lg.p-responsive.pt-6 > div > div:nth-child(2) > article:nth-child(n) > p")
		repos = append(repos, repo)
	})

	// Before finishing a scrape print "Finished ..."
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		createJsonFile("results.json", repos)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Visit the trending go page
	c.Visit("https://github.com/trending/go?since=daily")
}

// Function that runs everything
func main() {
	scrapeHandler()
}