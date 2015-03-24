package main

import (
	"github.com/gin-gonic/gin"
	scrape "github.com/tnngo2/scrape/lib"
	//"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	r := gin.Default()
	r.Use(Logger())
	r.LoadHTMLGlob("templates/*")

	r.GET("/topic", func(c *gin.Context) {
		c.Request.ParseForm()

		url := c.Request.Form.Get("url")
		message := scrape.GetWordList(url)
		c.String(http.StatusOK, message)
	})

	r.GET("/mean", func(c *gin.Context) {
		var response string
		const (
			BASE_URL = "http://dictionary.cambridge.org/search/british/direct/?q="
		)

		c.Request.ParseForm()

		wordList := c.Request.Form.Get("wordList")
		wordSlice := strings.Split(wordList, "\n")

		for i := range wordSlice {
			request := BASE_URL + url.QueryEscape(wordSlice[i])
			word := wordSlice[i]
			word = strings.TrimSpace(word)

			meaning, pronounce, example, guideword := scrape.GetWordMeaning(request)

			elem := word + "\t" + meaning + "\t" + pronounce + "\t" + example + "\t" + guideword + "\n"
			response = response + elem
		}

		obj := gin.H{"response": response}
		c.HTML(http.StatusOK, "mean.tmpl", obj)
	})

	r.Run(":" + os.Getenv("PORT"))
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
