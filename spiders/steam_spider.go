package spiders

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// SteamSpider -- mandate that SteamSpider implement BaseSpider + Crawler
type SteamSpider struct {
	BaseSpider
	Crawler
}

// Name -- steam
func (ss SteamSpider) Name() string {
	return "steam"
}

// BaseURL -- base URL for search Steam games
func (ss SteamSpider) BaseURL() string {
	return "https://store.steampowered.com/search/?term="
}

// Selector -- CSS selectors for finding the first result from searching
func (ss SteamSpider) Selector() string {
	return "#search_result_container > div > a:nth-of-type(1)"
}

func (ss SteamSpider) assembleSearchURL(searchTerm string) string {
	// trim spaces, then URL-encoded all spaces
	return fmt.Sprintf("%s%s", ss.BaseURL(), strings.ReplaceAll(strings.Trim(searchTerm, " "), " ", "%20"))
}

// Search -- initiate a search using the given searchTerm; return a JSON-encoded CrawlResult
func (ss SteamSpider) Search(searchTerm string) string {
	searchURL := ss.assembleSearchURL(searchTerm)
	fmt.Printf("Visiting %s\n", searchURL)

	// create a channel to return the CrawlResult
	resultChannel := make(chan string)

	Collector.OnHTML(ss.Selector(), func(e *colly.HTMLElement) {
		name := e.ChildText(".title")
		price, _ := strconv.Atoi(e.ChildAttr(".col.search_price_discount_combined", "data-price-final"))
		fmt.Printf("Result for '%s' -> {%s : %d}\n", searchTerm, name, price)

		crawlResult, err := json.Marshal(CrawlResult{name, price})
		if err != nil {
			resultChannel <- ""
			log.Fatal(err)
		}
		resultChannel <- string(crawlResult)
	})

	// Visit the assembled search URL
	Collector.Visit(searchURL)

	// wait for res channel to return
	return <-resultChannel
}
