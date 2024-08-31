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

	c.OnHTML("div[data-test-course-name]", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.Visit("https://app.codecrafters.io/catalog")

	fmt.Println("webscraper 1.0 ", c)
}
