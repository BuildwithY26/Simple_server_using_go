package main

import(
	"net/http"
	"time"
	"io"
)

type GoFeed struct {
	Channel struct{
		Title string `xml:"title"`
		Link string `xml:"link"`
		Description string `xml:"description"`
		Language string `xml:"language"`
		Item []GoItem `xml:"item"`
	}
}

type GoItem struct {
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`
	PubDate string `xml:"pubDate"`
}
func urlForFeed(url string)(GoFeed, error){
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := httpClient.Get(url)
	if err != nil {
		return GoFeed{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	goFeed := GoFeed{}
	err = xml.Unmarshal(dat, &goFeed)
	if err != nil {
		return GoFeed{}, err
	}
	return goFeed, nil
}