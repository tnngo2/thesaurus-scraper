package main

import (
	"github.com/gin-gonic/gin"
	scrape "github.com/tnngo2/scrape/lib"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()

	// This handler will match /user/john but will not match neither /user/ or /user
	r.GET("/topic", func(c *gin.Context) {
		c.Request.ParseForm()

		url := c.Request.Form.Get("url")
		message := scrape.GetWordList(url)
		c.String(http.StatusOK, message)
	})

	// Listen and server on 0.0.0.0:8080
	//r.Run(":80")
	r.Run(":" + os.Getenv("PORT"))
}
