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

	fmt.Println("webscraper 1.0 ", c)
}
