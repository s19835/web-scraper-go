package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type items struct {
	Name        string `json:"name"`
	Description string `json:"description"`
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
		}
		fmt.Println(courses)
	})

	c.Visit("https://app.codecrafters.io/catalog")

	fmt.Println("webscraper 1.0 ", c)
}
