package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
    // Instantiate default collector
    c := colly.NewCollector()

	c.OnHTML("a.result-image", func(e *colly.HTMLElement) {
		// collects all shopping item links
		link := e.Attr("href")
		// uses the text as a conditional/filter
		if (e.Text == "\n                $0\n        ") {
			// if the item is $0 then give us the link!
			fmt.Printf("$0 -> %q \n", link)
		}
	})

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })

    c.OnError(func(_ *colly.Response, err error) {
        fmt.Println("Something went wrong:", err)
    })

    c.OnResponse(func(r *colly.Response) {
        fmt.Println("Visited", r.Request.URL)
    })

    c.OnScraped(func(r *colly.Response) {
        fmt.Println("Finished", r.Request.URL)
    })

    // Start scraping on https://hackerspaces.org
    c.Visit("https://bakersfield.craigslist.org/d/for-sale/search/sss")
}
