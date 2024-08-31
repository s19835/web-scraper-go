package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type items struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgUrl"`
}

func main() {
	c := colly.NewCollector()

	// c.OnHTML("div[data-test-course-name]", func(h *colly.HTMLElement) {
	// 	fmt.Println(h.ChildText("span"))
	// })

	c.OnHTML("a[data-test-course-card]", func(h *colly.HTMLElement) {
		fmt.Println(h.ChildText("div[data-test-course-name]"), " : ", h.ChildText("div[data-test-course-description]"))
	})

	c.Visit("https://app.codecrafters.io/catalog")

	fmt.Println("webscraper 1.0 ", c)
}
