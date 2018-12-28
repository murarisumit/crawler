package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/murarisumit/crawler/pkg/scrawler"
)

func main() {
	log.Println(" Starting ..")
	hurl, _ := url.Parse("http://www.google.com")
	fmt.Printf("Host is: %s \n", hurl.Host)
	config := getConfig()
	crawler := scrawler.NewCrawler(config)
	log.Println("====== Starting crawling =====")
	crawler.Crawl(config.BaseURL.String(), config.Depth)
	log.Println("====== Done  crawling =====")
	crawler.Print()
}

func getConfig() scrawler.Config {
	cfg := scrawler.Config{}
	cfg.Concurrency = 2
	cfg.Depth = 2
	cfg.BaseURL, _ = url.Parse("https://monzo.com")
	return cfg
}
