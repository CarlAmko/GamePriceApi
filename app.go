package main

import (
	"cache"
	"fmt"
	"net/http"
	"spiders"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create all spiders
	createSpiders()

	r := gin.Default()
	r.GET("/crawl/:provider/:query", func(c *gin.Context) {
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

	// listen and serve
	r.Run()
}

func createSpiders() {
	spiders.Spiders["steam"] = spiders.SteamSpider{}
}
