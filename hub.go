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

type Repo struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
}

func createJsonFile(filename string, repos []Repo) {
	jsonData, err := json.MarshalIndent(repos, "", "    ")

	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(filename, jsonData, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func scrapeHandler() {
	c := colly.NewCollector()
	repos := make([]Repo, 0)

	// On every a element which has title attribute call callback
	c.OnHTML("body > div.application-main > main > div.explore-pjax-container.container-lg.p-responsive.pt-6 > div > div:nth-child(2) > article:nth-child(n)", func(e *colly.HTMLElement) {
		// TODO: Clean up data before assigning to struct
		repo := Repo{}
		info := e.ChildText("body > div.application-main > main > div.explore-pjax-container.container-lg.p-responsive.pt-6 > div > div:nth-child(2) > article:nth-child(n) > h1")
		name := strings.Replace(info, "\n\n\n     ", "", -1)
		repo.Name = name

		repo.Description = e.ChildText("body > div.application-main > main > div.explore-pjax-container.container-lg.p-responsive.pt-6 > div > div:nth-child(2) > article:nth-child(n) > p")
		repos = append(repos, repo)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		createJsonFile("results.json", repos)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://github.com/trending/go?since=daily")
}

func main() {
	scrapeHandler()
}

		// // usernameAndRepoName := strings.Split(repo.Name, "/")
		// // username := strings.TrimSpace(usernameAndRepoName[0])
		// // repoName := strings.TrimSpace(usernameAndRepoName[1])
		// // repo.Name = username + "/" + repoName
		// // cleanRepoName := strings.Replace(repo.Name, `\n\n\n\n`, "", -1)
		// repo.Name = strings.TrimSpace(e.ChildText("body > div.application-main > main > div.explore-pjax-container.container-lg.p-responsive.pt-6 > div > div:nth-child(2) > article:nth-child(n) > h1"))
		// // repo.Name = strings.Replace(e.Text, `\n\n\n`, "      ", -1)