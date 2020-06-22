package spiders

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// G2ASpider -- mandate that G2ASpider implement BaseSpider + Crawler
type G2ASpider struct {
	BaseSpider
	Crawler
}

// Name -- steam
func (gs G2ASpider) Name() string {
	return "g2a"
}

// BaseURL -- base URL for search Steam games
func (gs G2ASpider) BaseURL() string {
	return "https://www.g2a.com/en-us/search?query="
}

// Selector -- CSS selectors for finding the first result from searching
func (gs G2ASpider) Selector() string {
	return "div#app"
}

func (gs G2ASpider) assembleSearchURL(searchTerm string) string {
	// trim spaces, then URL-encoded all spaces
	return fmt.Sprintf("%s%s", gs.BaseURL(), strings.ReplaceAll(strings.Trim(searchTerm, " "), " ", "%20"))
}

// Search -- initiate a search using the given searchTerm; return a JSON-encoded CrawlResult
func (gs G2ASpider) Search(searchTerm string) string {
	// clone the collector for this crawl
	clone := Collector.Clone()

	searchURL := gs.assembleSearchURL(searchTerm)
	clone.OnRequest(func(r *colly.Request) {
		// notify visiting
		fmt.Printf("Visiting %s\n", searchURL)
	})

	// default to finding no results
	result := fmt.Sprintf("No result found for '%s'", searchTerm)

	clone.OnResponse(func (r * colly.Response) {
		fmt.Printf("Response body: %s.\n", string(r.Body))
	})

	clone.OnHTML(gs.Selector(), func(e *colly.HTMLElement) {
		name := e.ChildText("h3.Card__title > a")
		price, _ := strconv.Atoi(e.ChildText("span.Card__price-cost price"))
		fmt.Printf("Found result for '%s' -> {%s : %d}\n", searchTerm, name, price)

		crawlResult, err := json.Marshal(CrawlResult{gs.Name(), name, price})
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
