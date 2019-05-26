package spiders

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// HBSpider -- mandate that HBSpider implement BaseSpider + Crawler
type HBSpider struct {
	BaseSpider
	Crawler
}

// Name -- steam
func (hbs HBSpider) Name() string {
	return "hb"
}

// BaseURL -- base URL for search Steam games
func (hbs HBSpider) BaseURL() string {
	return "https://www.humblebundle.com/store/api/search?sort=alphabetical&filter=all&search="
}

// Selector -- CSS selectors for finding the first result from searching
func (hbs HBSpider) Selector() string {
	return "ul.entities-list.js-entities-list > li:nth-of-type(1)"
}

func (hbs HBSpider) assembleSearchURL(searchTerm string) string {
	// trim spaces, then URL-encoded all spaces
	return fmt.Sprintf("%s%s%s", hbs.BaseURL(), strings.ReplaceAll(strings.Trim(searchTerm, " "), " ", "%20"), "&request=1")
}

// Search -- initiate a search using the given searchTerm; return a JSON-encoded CrawlResult
func (hbs HBSpider) Search(searchTerm string) string {
	// clone the collector for this crawl
	clone := Collector.Clone()

	searchURL := hbs.assembleSearchURL(searchTerm)
	clone.OnRequest(func(r *colly.Request) {
		// notify visiting
		fmt.Printf("Visiting %s\n", searchURL)
	})

	// default to finding no results
	result := fmt.Sprintf("No result found for '%s'", searchTerm)

	clone.OnResponse(func(r *colly.Response) {
		jsonString := string(r.Body)
		name := gjson.Get(jsonString, "results.0.human_name").Str
		price := gjson.Get(jsonString, "results.0.current_price.0").Num

		if name != "" && price > 0.00 {
			res, _ := json.Marshal(CrawlResult{hbs.Name(), name, int(price * 100)})
			result = string(res)
		}
	})

	// Visit the assembled search URL
	clone.Visit(searchURL)
	clone.Wait()

	// get first result from channel
	fmt.Printf("Got result: %s\n", result)
	return result
}
