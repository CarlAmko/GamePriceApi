package cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

// Env -- parsed environment variables from .env file
var Env, _ = godotenv.Read()

// Create a new Redis client
var parsedConnectionOptions, _ = redis.ParseURL(Env["REDIS_URL"])
var redisClient = redis.NewClient(parsedConnectionOptions)

// Default cache time to 1 hour
var cacheTime, _ = time.ParseDuration("1h")

// Get -- fetch a cached search result from Colly, if it exists
func Get(provider string, searchTerm string) (string, error) {
	return redisClient.Get(fmt.Sprintf("%s:%s", provider, searchTerm)).Result()
}

// Set -- store a crawl result into Redis
func Set(provider string, searchTerm string, result string) (bool, error) {
	// for safety, only set the key if it does not already exist
	return redisClient.SetNX(fmt.Sprintf("%s:%s", provider, searchTerm), result, cacheTime).Result()
}
