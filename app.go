package main

import (
	"time"
	"encoding/json"
	"GamePriceApi/cache"
	"GamePriceApi/spiders"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create all spiders
	createSpiders()

	r := gin.Default()
	r.GET("/crawl/:query/:provider", func(c *gin.Context) {
		provider := strings.ToLower(c.Param("provider"))
		query := c.Param("query")

		// find spider for given provider
		s := spiders.Spiders[provider]
		if s == nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Invalid provider: %s.", provider))
			return
		}

		// first check if this result is cached
		res, err := cache.Get(provider, query)

		// if err is returned, then this value is not cached
		if err != nil {
			// run a search with given query
			json := s.Search(query)
			// return JSON string result
			c.JSON(http.StatusOK, json)

			// cache result
			cache.Set(provider, query, json)
		} else {
			// value is cached, just return it
			c.JSON(http.StatusOK, res)
		}
	})

	r.GET("/crawl/:query", func(c *gin.Context) {
		query := c.Param("query")

		// map all responses
		results := make([]string, 0)

		// query all spiders
		for provider, spider := range spiders.Spiders {
			// first check if this result is cached
			cachedResult, err := cache.Get(provider, query)

			// if err is returned, then this value is not cached
			if err != nil {
				// run a search with given query
				cachedResult = spider.Search(query)

				// cache result
				cache.Set(provider, query, cachedResult)
			}

			// filter out non-results
			if !strings.Contains(strings.ToLower(cachedResult), "no result found") {
				results = append(results, cachedResult)
			}
		}
		fmt.Printf("%s\n", results)
		// marshall results
		resp := spiders.QueryAllResult{results, time.Now().Unix()}

		jsonResult, _ := json.Marshal(resp)
		fmt.Printf("%s\n", jsonResult)

		c.JSON(http.StatusOK, string(jsonResult))
	})

	// listen and serve
	r.Run()
}

func createSpiders() {
	spiders.Spiders["steam"] = spiders.SteamSpider{}
	spiders.Spiders["g2a"] = spiders.G2ASpider{}
	spiders.Spiders["hb"] = spiders.HBSpider{}
}
