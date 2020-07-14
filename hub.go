package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
	"log"
	"github.com/gocolly/colly"
)

// TYPE
type Repo struct {
	Name 	      string  
	Description   string  
	Stars         string
	Users         string
}

// TYPE
type Repos struct {
	Repos string
}

func jsonRepo(filename string, e Repos) {
	jsonRepo, err := json.Marshal(e)

	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("output.json", jsonRepo, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func main() {
	c := colly.NewCollector()
	var repos [] Repos

	// On every a element which has title attribute call callback
	c.OnHTML("body > div.application-main > main > div.explore-pjax-container.container-lg.p-responsive.pt-6 > div > div:nth-child(2) > article:nth-child(n)", func(e *colly.HTMLElement) {
		// e.ForEach("body > div.application-main > main > div.explore-pjax-container.container-lg.p-responsive.pt-6 > div > div:nth-child(2) > article:nth-child(n) > h1", func(e *colly.HTMLElement) {
		// })

		// e.ForEach("body > div.application-main > main > div.explore-pjax-container.container-lg.p-responsive.pt-6 > div > div:nth-child(2) > article:nth-child(n) > p", func(e *colly.HTMLElement) {
		// })
		repo := Repos{Repos: e.Text}

		j, _ := json.Marshal(repo)
		_ = ioutil.WriteFile("output.json", j, os.ModePerm)

		for _, repo := range repos {
			jsonRepo("output.json", repo)
		}

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://github.com/trending/go?since=daily")

}
