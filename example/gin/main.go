package main

import (
	"embed"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/simbafs/kama"
)

//go:embed all:static
var static embed.FS

func main() {
	r := gin.Default()

	k := kama.New(static)

	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from the backend",
		})
	})

	r.Use(k.Gin())

	log.Fatal(r.Run(":3000"))
}
