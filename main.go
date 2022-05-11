package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

// Config model
type Config struct {
	City     string   `json:"city"`
	Area     string   `json:"area"`
	Category string   `json:"category"`
	Queries  []string `json:"queries"`
}

// Default error handler
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Read in the config file
	var config Config
	data, err := os.ReadFile("config.json")
	checkError(err)

	err = json.Unmarshal(data, &config)
	checkError(err)

	// Append slash to the end of Area for url creation
	if config.Area != "" {
		config.Area += "/"
	}

	c := colly.NewCollector(
		colly.AllowedDomains("craigslist.org", fmt.Sprintf("%v.craigslist.org", config.City)),
	)

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	listings := make([]string, 0)

	// Scrape different categories
	switch config.Category {
	case "sss":
		// On every a element which has href attribute call callback

		url := fmt.Sprintf("https://%v.craigslist.org/search/%v%v?postedToday=1", config.City, config.Area, config.Category)

		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if link == "" {
				return
			}

			switch e.Attr("class") {
			case "result-image gallery":
				fmt.Printf("Link found: %v\n", link)
				listings = append(listings, link)
			case "button next":
				c.Visit(e.Request.AbsoluteURL(link))
			}
		})

		// Start scraping on url
		c.Visit(url)

		fmt.Println(len(listings))
	default:
		fmt.Printf("Category %v is not supported\n", config.Category)
	}
}
