package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// On every a element which has title attribute call callback
	c.OnHTML("body > div.application-main > main > div.explore-pjax-container.container-lg.p-responsive.pt-6 > div > div:nth-child(2) > article:nth-child(n) > h1", func(e *colly.HTMLElement) {
		fmt.Printf("User/Name: %s\n", e.Text)
	})

	// c.OnHTML("body > div.application-main > main > div.explore-pjax-container.container-lg.p-responsive.pt-6 > div > div:nth-child(2) > article:nth-child(n) > p", func(e *colly.HTMLElement) {
	// 	fmt.Printf("Description: %s\n", e.Text)
	// })

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://github.com/trending/go?since=daily")

}
