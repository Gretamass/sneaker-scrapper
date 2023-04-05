package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/http"
	"os"
)

type item struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/nike", kixartNike)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func kixartNike(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()
	var items []item

	c.OnHTML("a[class=pr_productItem]", func(element *colly.HTMLElement) {
		item := item{
			Name:  element.ChildText("div.bottom div.title"),
			Price: element.ChildText("div.bottom div.price"),
		}
		items = append(items, item)
	})

	c.OnHTML("a.next", func(element *colly.HTMLElement) {
		nextPage := element.Request.AbsoluteURL(element.Attr("href"))
		c.Visit(nextPage)
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println(request.URL.String())
	})

	c.Visit("https://www.kixart.lt/kedai-laisvalaikiui?vars/manufacturer/1")

	content, err := json.MarshalIndent(items, "", "    ")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile("KixartNikeProducts.json", content, 0644)
	if err != nil {
		fmt.Fprintf(w, "Error writing file: ")
	}
	json.NewEncoder(w).Encode(content[0])
}
