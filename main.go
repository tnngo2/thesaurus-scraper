package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	scrape "github.com/tnngo2/scrape/lib"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	//r := gin.New()
	r.Use(Logger())
	r.LoadHTMLGlob("templates/*")

	// This handler will match /user/john but will not match neither /user/ or /user
	r.GET("/topic", func(c *gin.Context) {
		c.Request.ParseForm()

		url := c.Request.Form.Get("url")
		message := scrape.GetWordList(url)
		c.String(http.StatusOK, message)
	})

	r.GET("/mean", func(c *gin.Context) {
		url := "http://dictionary.cambridge.org/search/british/direct/?q="
		c.Request.ParseForm()
		wordList := c.Request.Form.Get("wordList")
		log.Println("....")
		if wordList != "" {
			url = url + wordList
			word := wordList
			meaning, pronounce, example, guideword := scrape.GetWordMeaning(url)
			elem := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\n", word, meaning, pronounce, example, guideword)
			obj := gin.H{"title": elem}
			c.HTML(http.StatusOK, "mean.tmpl", obj)
		}
	})

	// Listen and server on 0.0.0.0:8080
	//r.Run(":80")
	r.Run(":" + os.Getenv("PORT"))
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
