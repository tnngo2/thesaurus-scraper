package main

import (
	"github.com/gin-gonic/gin"
	scrape "github.com/tnngo2/scrape/lib"
	"net/http"
	//"os"
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

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/join/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Params.ByName("name")
		action := c.Params.ByName("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	r.GET("/vn/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		message := "vn " + name
		c.String(http.StatusOK, message)
	})

	// Listen and server on 0.0.0.0:8080
	//r.Run(":80")
	r.Run("3001")
}
