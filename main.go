package main

import (
	"fmt"
	"log"
	"github.com/gocolly/colly"
)

type Quote struct {
	Text string
	Author string
}

func main(){
	var quotes []Quote
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)

	c.OnError(func(r *colly.Response, err error) {
		log.Fatal("Request failed:", r.StatusCode, err)
	})

	c.OnHTML(".quote", func(e *colly.HTMLElement) {
		quote := Quote{
			Text: e.ChildText(".text"),
			Author: e.ChildText(".author"),
		}
		quotes = append(quotes, quote)
	})

	err := c.Visit("https://quotes.toscrape.com/tag/inspirational/")
	if err != nil {
		log.Fatal("Visit failed:", err)
	}

	for idx, q := range quotes {
		fmt.Printf("%d, %s\n - %s\n\n", idx+1, q.Text, q.Author)
	}


}