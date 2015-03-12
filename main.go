package main

import (
	"fmt"
	scrape "github.com/tnngo2/scrape/lib"
	"net/http"
	"os"
)

func main() {
	//result := scrape.ImportWordList("../list.md")
	//scrape.PrintSlice(result)

	//url := "http://dictionary.cambridge.org/dictionary/british/clamp"
	//result := GetThesaurusUrl(url)
	//PrintSlice(result)
	//GetWordList(url)

	http.HandleFunc("/", Thesaurus)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func Thesaurus(res http.ResponseWriter, req *http.Request) {
	url := req.URL.Query()["u"]
	result := scrape.GetThesaurusUrl(url[0])
	fmt.Fprintln(res, scrape.PrintSliceHtml(result))
}
