package scrape

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func GetThesaurusUrl(url string) []string {
	var href []string
	doc := GetNewDocument(url)

	mainTopic, _ := doc.Find("#cdo-right-shaded #cdo-smartt #cdo-main-cloud-container a.see-more").Attr("href")
	href = append(href, mainTopic)

	doc.Find("#cdo-other-topics a").Each(func(i int, s *goquery.Selection) {
		otherTopic, _ := s.Attr("href")
		if otherTopic != "" {
			href = append(href, otherTopic)
		}
	})

	return href
}

func GetWordList(url string) string {
	result := ""
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#cdo-smartt #cdo-main-cloud .cdo-cloud-content ul li").Each(func(i int, s *goquery.Selection) {
		word := s.Find("b").Text()
		href, _ := s.Find("a").Attr("href")
		if href != "" {
			meaning, pronounce, example, guideword := GetWordMeaning(href)
			result += fmt.Sprintf("%s\t%s\t%s\t%s\t%s\n", word, meaning, pronounce, example, guideword)
		}
	})

	return result
}

func GetWordMeaning(url string) (string, string, string, string) {
	doc := GetNewDocument(url)

	meaning := doc.Find(".def").First().Text()
	pronounce := doc.Find("span.us span.pron").First().Text()
	example := doc.Find(".examp").First().Text()
	guideword := doc.Find(".guideword").First().Text()
	guideword = strings.TrimSpace(guideword)

	guideword = doc.Find(".posgram").First().Text() + " " + guideword
	guideword = strings.Replace(guideword, "\n", "", -1)
	guideword = strings.Replace(guideword, "\t", "", -1)

	if pronounce != "" {
		pronounce = doc.Find("#entryContent span.uk span.pron").Text()
	}

	return meaning, pronounce, example, guideword
}

func GetNewDocument(url string) *goquery.Document {
	doc, err := goquery.NewDocument(strings.TrimSpace(url))
	if err != nil {
		log.Fatal(err)
	}
	return doc
}
