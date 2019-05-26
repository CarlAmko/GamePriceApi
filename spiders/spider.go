package spiders

import "github.com/gocolly/colly"

// Crawler -- interface defining actions that Spiders can perform
type Crawler interface {
	assembleSearchURL(searchTerm string) string
	Search(searchTerm string) string
}

// BaseSpider -- Base type declaration for any spider
type BaseSpider interface {
	Name() string
	BaseURL() string
	Selector() string
}

// CrawlResult -- Return type resulting from a crawl
type CrawlResult struct {
	Provider string `json:"provider"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
}

type QueryAllResult struct {
	Results []string `json:"results"`
	Time    int64    `json:"query_time"`
}

// Collector -- Singleton, reusable, asynchronous collector
var Collector = colly.NewCollector(colly.Async(true))

// Spiders -- Map of all providers to their spider implementations
var Spiders = make(map[string]Crawler)
