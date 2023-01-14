package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("RDS_ADDR"),
	Username: os.Getenv("RDS_USER"),
	Password: os.Getenv("RDS_PWD"),
	DB:       0, // use default DB
})

// TODO: handle duplicated longUrl
// TODO: handle encoding method
// TODO: load balancing

func encode(key uint64) string {
	element := "0123456789abcdefghijklmnopqustuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	shorturl := ""
	for key > 0 {
		i := key % 62
		key = key / 62
		shorturl += string(element[i])
	}
	return shorturl
}

func shorten(c *gin.Context) {
	longUrl := c.Query("longUrl")

	if len(longUrl) == 0 {
		c.JSON(400, gin.H{
			"msg": "incorrect query",
		})
		return
	}

	shortUrl := encode(rand.Uint64())
	err := rdb.Set(ctx, shortUrl, longUrl, 0).Err()
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "error",
		})
		panic(err)
	}

	fmt.Println("[DEBUG]", shortUrl, longUrl)
	c.JSON(200, gin.H{
		"shortUrl": shortUrl,
		// "longUrl":  longUrl,
	})
}

func expand(c *gin.Context) {
	key := c.Query("shortUrl")

	if len(key) == 0 {
		c.JSON(400, gin.H{
			"msg": "incorrect query",
		})
		return
	}

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "cannot find the query shorturl",
		})
		panic(err)
	}

	fmt.Println("[DEBUG]", key, val)
	c.JSON(200, gin.H{
		// "shortUrl": key,
		"longUrl": val,
	})
}

func main() {
	r := gin.Default()
	r.GET("/api/shorten", shorten)
	r.GET("/api/expand", expand)

	fmt.Println("server starting on port 8080")
	r.Run(":8080")
}
