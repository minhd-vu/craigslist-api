package main

import (
	"encoding/json"
	"fmt"
	"os"
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

	// Scrape different categories
	switch config.Category {
	case "sss":
		url := fmt.Sprintf("https://%v.craigslist.org/search/%v%v?postedToday=1", config.City, config.Area, config.Category)
		fmt.Println(url)
	default:
		fmt.Printf("Category %v is not supported\n", config.Category)
	}
}
