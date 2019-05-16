package main

import (
	"strings"
	"time"
	"net/http"
	"fmt"
	"github.com/mmcdole/gofeed"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/jdkato/prose.v2"
)

func feed(i int) (url string){

	fp := gofeed.NewParser()
	
	feed, err := fp.ParseURL("http://feeds.reuters.com/reuters/UKTopNews")
	
	if err != nil {
		panic(err)
	}

	fmt.Println(feed.Items[i].Title)

	url = fmt.Sprintf(feed.Items[i].Link)

	return url

}

func getContent(i int)(content string){
	
	url := feed(i)

	req, err := http.NewRequest("GET",url, nil)
		if err != nil {}

  	client := &http.Client{Timeout: time.Second * 10}
	
	resp, err := client.Do(req)
	if err != nil { fmt.Println("req err")}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
  	if err != nil {}

	s:= fmt.Sprintf(doc.Find(".StandardArticleBody_container").Text())
	
	return s
}

func wordCount(str string) map[string]int {
    wordList := strings.Fields(str)
    counts := make(map[string]int)
    for _, word := range wordList {
        _, ok := counts[word]
        if ok {
            counts[word] += 1
        } else {
            counts[word] = 1
        }
    }
    return counts
}

var input string

func ent(c chan string){

	s:= getContent(0)

	doc, err := prose.NewDocument(s)
	if err != nil {}

	// Iterate over the doc's named-entities:
    for _, ent := range doc.Entities() {
		c <- ent.Text
	}
}


var fields []string

func analysis(c chan string){ 
for{	
   input := <- c
   fields:= append(fields, input)
   fmt.Println("=>",fields)	
}
}


func main() {

	var c chan string = make(chan string)

	go ent(c)
	go analysis(c)

	var input string
   fmt.Scanln(&input)

}