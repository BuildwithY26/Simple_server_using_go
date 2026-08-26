package main

import (
	"context"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/BuildwithY26/go/internal/database"
)

func startScraping(db *database.Queries, 
	concurrency int, timeBwRequests time.Duration) {
		log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBwRequests)
		ticker := time.NewTicker(timeBwRequests)
		for ; ; <-ticker.C{
			feeds, err := db.GetNextFeedsToFetch(
				context.Background(),
				int32(concurrency),
			)
			if err != nil{
				log.Println("Error fetching feeds:", err)
				continue
			}
			log.Println("Found %v feeds to fetch", len(feeds))

			wg := &sync.WaitGroup{}
			for _, feed := range feeds {
				wg.Add(1)

				go scrapeFeed(db, wg, feed)
			}
			wg.Wait()
		}
}


func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed)  {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil{
				log.Println("Couldn't mark feed %s as fetched %v", feed.Name, err)
				return
			}

			goFeed, err := urlToFeed(feed.Url)
			if err != nil{
				log.Println("Error fetching feed %s: %v", feed.Name, err)
				return
			}
			for _, item := range goFeed.Channel.Item {
		log.Println("Found post", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(goFeed.Channel.Item))
}

type goFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []GoItem `xml:"item"`
	} `xml:"channel"`
}

func fetchFeed(feedURL string) (*goFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := httpClient.Get(feedURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var goFeed GoFeed
	err = xml.Unmarshal(dat, &goFeed)
	if err != nil {
		return GoFeed{}, err
	}

	return &goFeed, nil
}
}
