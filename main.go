package main

import (
	"log"
	"net/url"

	"github.com/murarisumit/crawler/pkg/scrawler"
)

type Crawler interface {
	Crawl(string, int)
	Print()
}

func main() {
	log.Println(" Starting ..")
	config := getConfig()
	var crawler Crawler = scrawler.NewCrawler(config)
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
	cfg.ExcludedPath = []string{
		"/cdn-cgi",
		"/legal",
		"/static",
		"/blog",
	}
	cfg.ExcludedSubdomain = []string{
		"www.monzo.com",
	}

	return cfg
}
