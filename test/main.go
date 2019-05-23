package main

import (
	"context"
	//"strings"
	"time"
	"net/http"
	"fmt"
	"github.com/mmcdole/gofeed"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/jdkato/prose.v2"
	language "cloud.google.com/go/language/apiv1"
  languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"

)

func feed(i int) (url string, title string, pubDate string){

	fp := gofeed.NewParser()
	
	feed, err := fp.ParseURL("https://www.economist.com/business/rss.xml")
	
	if err != nil {
		panic(err)
	}

	title = fmt.Sprintf(feed.Items[i].Title)

	url = fmt.Sprintf(feed.Items[i].Link)

	pubDate = fmt.Sprintf(feed.Items[i].Published)

	return url, title, pubDate

}

func getContent(i int)(content string, url string, title string, pubDate string){
	
	url, title, pubDate = feed(i)

	req, err := http.NewRequest("GET",url, nil)
		if err != nil {}

  	client := &http.Client{Timeout: time.Second * 10}
	
	resp, err := client.Do(req)
	if err != nil { fmt.Println("req err")}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
  	if err != nil {}

	s:= fmt.Sprintf(doc.Find(".blog-post__text").Text())

//	s:= fmt.Sprintf(doc.Find(".StandardArticleBody_container").Text())
	
	fmt.Println(s)
	
	return s, url, title, pubDate
}

func analyzeEntities(ctx context.Context, client *language.Client, text string) (*languagepb.AnalyzeEntitiesResponse, error) {
	return client.AnalyzeEntities(ctx, &languagepb.AnalyzeEntitiesRequest{
					Document: &languagepb.Document{
									Source: &languagepb.Document_Content{
													Content: text,
									},
									Type: languagepb.Document_PLAIN_TEXT,
					},
					EncodingType: languagepb.EncodingType_UTF8,
	})
}

func analyse(i int){

	text, url, title, pubDate := getContent(i)

		fmt.Println()
		fmt.Println()
		fmt.Println(title)
		fmt.Println(pubDate)
		fmt.Println(url)

	doc, err := prose.NewDocument(text)
	if err != nil {}
	
	// Iterate over the doc's named-entities:
		ctx := context.Background()

		// Creates a client.
	
	for _, ent := range doc.Sentences() {
		
		text := fmt.Sprintf(ent.Text)

    client, err := language.NewClient(ctx)
    if err != nil {}
		
		analysis, err:= analyzeEntities(ctx, client, text)
		if err !=nil{}

		fmt.Println(len(analysis.Entities))

		for _, t := range analysis.Entities {
			if t.Type == 11 {
				fmt.Println(text)
				fmt.Println("[Time: =========>]",t.Name)
			}
		}
	}
}

func main(){
	for i:=0; i <20 ; i++ {
		analyse(i)
	}
}