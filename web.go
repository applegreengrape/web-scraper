package main

import (
	//"io/ioutil"
	"time"
	"net/http"
	"golang.org/x/net/html"
	"fmt"
	"github.com/mmcdole/gofeed"
)

func getHref(t html.Token) (ok bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}
	return
}

func main(){

	fp := gofeed.NewParser()
	
	feed, err := fp.ParseURL("http://feeds.reuters.com/reuters/UKTopNews")
	
	if err != nil {
		panic(err)
	}

	fmt.Println(feed.Items[0].Title)

	fmt.Println(feed.Items[0].Link)

	fmt.Println(len(feed.Items))

  req, err := http.NewRequest("GET", feed.Items[0].Link, nil)
	if err != nil {}

  client := &http.Client{Timeout: time.Second * 10}
	
	resp, err := client.Do(req)
	if err != nil { fmt.Println("req err")}

	b := resp.Body
	defer b.Close()

	z := html.NewTokenizer(b)
	fmt.Println("Got the tokenizer")

	for {
		tt := z.Next()
		testt := z.Token()

		switch {
		case tt == html.ErrorToken:

			return
		case tt == html.StartTagToken:
			isAnchor := testt.Data == "a"

			if !isAnchor {
				continue
			}

			ok, url := getHref(testt)
			if !ok {
				continue
			}

			fmt.Println("***", url)
		}
	}

}