package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func GetThesaurusUrl(url string) string {
	doc := GetNewDocument(url)
	href, _ := doc.Find("#cdo-right-shaded #cdo-smartt #cdo-main-cloud-container a.see-more").Attr("href")
	return href
}

func GetWordList(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#cdo-smartt #cdo-main-cloud .cdo-cloud-content ul li").Each(func(i int, s *goquery.Selection) {
		word := s.Find("b").Text()
		href, _ := s.Find("a").Attr("href")
		//fmt.Printf("%s\t%s \n", word, href)
		if href != "" {
			meaning, pronounce, example := GetWordMeaning(href)
			fmt.Printf("%s\t%s\t%s\t%s\n", word, meaning, pronounce, example)
		}
	})
}

func GetWordMeaning(url string) (string, string, string) {
	doc := GetNewDocument(url)

	meaning := doc.Find("#entryContent #1-1 span.def").Text()
	pronounce := doc.Find("#entryContent span.us span.pron").Text()
	example := doc.Find("#entryContent #1-1 .examp").First().Text()

	return meaning, pronounce, example
}

func GetNewDocument(url string) *goquery.Document {
	doc, err := goquery.NewDocument(strings.TrimSpace(url))
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func main() {
	url := GetThesaurusUrl("http://dictionary.cambridge.org/dictionary/british/deep")
	GetWordList(url)
}
