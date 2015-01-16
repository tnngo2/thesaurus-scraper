package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func GetThesaurusUrl() string {
	doc, err := goquery.NewDocument("http://dictionary.cambridge.org/dictionary/british/deep")
	if err != nil {
		log.Fatal(err)
	}

	// doc.Find("#cdo-right-shaded #cdo-smartt #cdo-main-cloud-container a.see-more").Each(func(i int, s *goquery.Selection) {
	// 	band := s.Find("h3").Text()
	// 	title := s.Find("i").Text()
	// 	fmt.Printf("Review %d: %s - %s\n", i, band, title)
	// })
	href, _ := doc.Find("#cdo-right-shaded #cdo-smartt #cdo-main-cloud-container a.see-more").Attr("href")

	fmt.Println("href : ", href)
	return href
}

func GetWordList(url string) {
	doc, err := goquery.NewDocument("http://dictionary.cambridge.org/dictionary/british/deep")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#cdo-smartt #cdo-main-cloud .cdo-cloud-content ul li").Each(func(i int, s *goquery.Selection) {
		word := s.Find("b").Text()
		def, _ := s.Find("a").Attr("href")
		fmt.Printf("%s : %s \n", word, def)
	})
}

func main() {
	url := GetThesaurusUrl()
	GetWordList(url)
}
