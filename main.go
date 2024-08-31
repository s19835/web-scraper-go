package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type items struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stages      string `json:"stages"`
}

func main() {
	c := colly.NewCollector()

	// c.OnHTML("div[data-test-course-name]", func(h *colly.HTMLElement) {
	// 	fmt.Println(h.ChildText("span"))
	// })

	c.OnHTML("a[data-test-course-card]", func(h *colly.HTMLElement) {
		courses := items{
			Name:        h.ChildText("div[data-test-course-name]"),
			Description: h.ChildText("div[data-test-course-description]"),
			Stages:      h.ChildText("span[class='text-xs text-gray-500']"),
		}
		fmt.Println(courses)
	})

	c.Visit("https://app.codecrafters.io/catalog")

	fmt.Println("webscraper 1.0 ", c)
}
