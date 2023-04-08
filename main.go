package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type item struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

func main() {
	// run automatedScrapers once immediately
	automatedScrapers()

	// run automatedScrapers every 6 hours
	ticker := time.NewTicker(6 * time.Hour)
	for range ticker.C {
		automatedScrapers()
	}
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func manualScrapers() {
	// KIXART
	//http.HandleFunc("/kixart/nike", kixartNike)
	//http.HandleFunc("/kixart/adidas", kixartAdidas)
	//http.HandleFunc("/kixart/puma", kixartPuma)
	//http.HandleFunc("/kixart/converse", kixartConverse)
	//http.HandleFunc("/kixart/air-jordan", kixartAirJordan)
	//http.HandleFunc("/kixart/new-balance", kixartNewBalance)
	//http.HandleFunc("/kixart/reebok", kixartReebok)
	//
	//// BALLZY
	//http.HandleFunc("/ballzy/nike", ballzyNike)
	//http.HandleFunc("/ballzy/converse", ballzyConverse)
	//
	//http.HandleFunc("/sizeer/nike", sizeerNike)
	//http.HandleFunc("/sil/nike", silNike)
	//http.HandleFunc("/snkrs/nike", snkrsNike)
}

func automatedScrapers() {
	kixartNike()
	kixartAdidas()
	kixartPuma()
	kixartConverse()
	kixartAirJordan()
	kixartNewBalance()
	kixartReebok()

	ballzyNike()
	ballzyConverse()

	sizeerNike()

	silNike()

	snkrsNike()
}

// KIXART
func kixartNike() {
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
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Response received:", response.StatusCode)
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.Visit("https://www.kixart.lt/kedai-laisvalaikiui?vars/manufacturer/1")

	content, err := json.MarshalIndent(items, "", "    ")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile("KixartNikeProducts.json", content, 0644)
}

func kixartAdidas() {
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
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Response received:", response.StatusCode)
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.Visit("https://www.kixart.lt/kedai-laisvalaikiui?vars/manufacturer/2")

	content, err := json.MarshalIndent(items, "", "    ")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile("KixartAdidasProducts.json", content, 0644)
}

func kixartPuma() {
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
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Response received:", response.StatusCode)
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.Visit("https://www.kixart.lt/kedai-laisvalaikiui?vars/manufacturer/9")

	content, err := json.MarshalIndent(items, "", "    ")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile("KixartPumaProducts.json", content, 0644)
}

func kixartConverse() {
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
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Response received:", response.StatusCode)
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.Visit("https://www.kixart.lt/kedai-laisvalaikiui?vars/manufacturer/5")

	content, err := json.MarshalIndent(items, "", "    ")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile("KixartConverseProducts.json", content, 0644)
}

func kixartAirJordan() {
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
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Response received:", response.StatusCode)
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.Visit("https://www.kixart.lt/kedai-laisvalaikiui?vars/manufacturer/4")

	content, err := json.MarshalIndent(items, "", "    ")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile("KixartAirJordanProducts.json", content, 0644)
}

func kixartNewBalance() {
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
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Response received:", response.StatusCode)
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.Visit("https://www.kixart.lt/kedai-laisvalaikiui?vars/manufacturer/13")

	content, err := json.MarshalIndent(items, "", "    ")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile("KixartNewBalanceProducts.json", content, 0644)
}

func kixartReebok() {
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
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Response received:", response.StatusCode)
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.Visit("https://www.kixart.lt/kedai-laisvalaikiui?vars/manufacturer/10")

	content, err := json.MarshalIndent(items, "", "    ")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile("KixartReebokProducts.json", content, 0644)
}

// BALLZY
func ballzyNike() {
	c := colly.NewCollector()
	var items []item

	c.OnHTML("div.product-item-details", func(element *colly.HTMLElement) {
		item := item{
			Name:  element.ChildText("strong.product-item-name a.product-item-link"),
			Price: element.ChildText("div.price-final_price span.normal-price span.price"),
		}
		items = append(items, item)
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println(request.URL.String())
	})
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Response received:", response.StatusCode)
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	// Scrape multiple pages by appending &p=2, &p=3, etc.
	// TODO: make this dynamic
	for i := 1; i <= 8; i++ {
		url := fmt.Sprintf("https://ballzy.eu/en/shoes.html?brands=7&p=%d", i)
		c.Visit(url)
	}

	content, err := json.MarshalIndent(items, "", "    ")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile("BallzyNikeProducts.json", content, 0644)
}

func ballzyConverse() {
	c := colly.NewCollector()
	var items []item

	c.OnHTML("div.product-item-details", func(element *colly.HTMLElement) {
		item := item{
			Name:  element.ChildText("strong.product-item-name a.product-item-link"),
			Price: element.ChildText("div.price-final_price span.normal-price span.price"),
		}
		items = append(items, item)
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println(request.URL.String())
	})
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Response received:", response.StatusCode)
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	// Scrape multiple pages by appending &p=2, &p=3, etc.
	// TODO: make this dynamic
	for i := 1; i <= 8; i++ {
		url := fmt.Sprintf("https://ballzy.eu/en/shoes.html?brands=203&p=%d", i)
		c.Visit(url)
	}

	content, err := json.MarshalIndent(items, "", "    ")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile("BallzyConverseProducts.json", content, 0644)
}

// Sizeer
func sizeerNike() {
	c := colly.NewCollector()
	var items []item

	c.OnHTML("div.b-itemList_desc", func(element *colly.HTMLElement) {
		item := item{
			Name:  element.ChildText("div.b-itemList_name a.b-itemList_nameLink"),
			Price: element.ChildText("div.b-itemList_prices p.b-itemList_price"),
		}
		items = append(items, item)
	})

	c.OnHTML("a.m-pagination_next", func(element *colly.HTMLElement) {
		nextPage := element.Request.AbsoluteURL(element.Attr("href"))
		c.Visit(nextPage)
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println(request.URL.String())
	})
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Response received:", response.StatusCode)
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.Visit("https://sizeer.lt/gamintojai/nike")

	content, err := json.MarshalIndent(items, "", "    ")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile("SizeerNikeProducts.json", content, 0644)
}

// SIL
func silNike() {
	c := colly.NewCollector()
	var items []item

	c.OnHTML("div.product-item-info", func(element *colly.HTMLElement) {
		var price string
		price = element.ChildText("span.special-price span.price")
		if price == "" {
			price = element.ChildText("span.price-final_price span.price")
		}
		item := item{
			Name:  element.ChildText("h2.product-item-name"),
			Price: price,
		}
		items = append(items, item)
	})

	c.OnHTML("a.action.next", func(element *colly.HTMLElement) {
		nextPage := element.Request.AbsoluteURL(element.Attr("href"))
		c.Visit(nextPage)
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println(request.URL.String())
	})
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Response received:", response.StatusCode)
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.Visit("https://www.sil.lt/avalyne?manufacturer=Nike")

	content, err := json.MarshalIndent(items, "", "    ")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile("SilNikeProducts.json", content, 0644)
}

// SNKRS
func snkrsNike() {
	c := colly.NewCollector()
	var items []item

	c.OnHTML("div.product-card-wrapper", func(element *colly.HTMLElement) {
		name := element.ChildText("h3.card__heading.h5 a.full-unstyled-link")
		codeIndex := strings.LastIndex(name, " ")
		if codeIndex > 0 {
			name = name[:codeIndex]
		}
		item := item{
			Name:  name,
			Price: element.ChildText("div.price__regular span.price-item--regular"),
		}
		items = append(items, item)
	})

	c.OnHTML("a.pagination__item--prev", func(element *colly.HTMLElement) {
		nextPage := element.Request.AbsoluteURL(element.Attr("href"))
		c.Visit(nextPage)
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println(request.URL.String())
	})
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Response received:", response.StatusCode)
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.Visit("https://snkrs.lt/collections/nike?filter.v.availability=1")

	content, err := json.MarshalIndent(items, "", "    ")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile("SnkrsNikeProducts.json", content, 0644)
}
