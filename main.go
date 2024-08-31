package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type items struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stages      string `json:"stages"`
}

func main() {
	c := colly.NewCollector()

	var courses []items

	// c.OnHTML("div[data-test-course-name]", func(h *colly.HTMLElement) {
	// 	fmt.Println(h.ChildText("span"))
	// })

	c.OnHTML("a[data-test-course-card]", func(h *colly.HTMLElement) {
		course := items{
			Name:        h.ChildText("div[data-test-course-name]"),
			Description: h.ChildText("div[data-test-course-description]"),
			Stages:      h.ChildText("span[class='text-xs text-gray-500']"),
		}

		courses = append(courses, course)
	})

	c.Visit("https://app.codecrafters.io/catalog")

	content, err := json.Marshal(courses)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("codecrafters-catalog.json", content, 0644)

	fmt.Println("webscraper 1.0 ", c)
}
