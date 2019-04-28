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
	// clone the collector for this crawl
	clone := Collector.Clone()

	searchURL := ss.assembleSearchURL(searchTerm)
	clone.OnRequest(func(r *colly.Request) {
		// notify visiting
		fmt.Printf("Visiting %s\n", searchURL)
	})

	// default to finding no results
	result := fmt.Sprintf("No result found for '%s'", searchTerm)

	clone.OnHTML(ss.Selector(), func(e *colly.HTMLElement) {
		name := e.ChildText(".title")
		price, _ := strconv.Atoi(e.ChildAttr(".col.search_price_discount_combined", "data-price-final"))
		fmt.Printf("Found result for '%s' -> {%s : %d}\n", searchTerm, name, price)

		crawlResult, err := json.Marshal(CrawlResult{name, price})
		if err != nil {
			// ensure that the channel doesn't block forever if JSON parsing errors
			result = fmt.Sprintf("Could not parse response for '%s'", searchTerm)
			log.Fatal(err)
		} else {
			result = string(crawlResult)
		}
	})

	// Visit the assembled search URL
	clone.Visit(searchURL)
	clone.Wait()

	// get first result from channel
	fmt.Printf("Got result: %s\n", result)
	return result
}
