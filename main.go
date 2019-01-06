package main

import (
	"log"
	"net/url"

	"github.com/murarisumit/crawler/pkg/crawler"
	"github.com/murarisumit/crawler/pkg/web"
)

type Crawler interface {
	Crawl(string, int)
}

type Website interface {
	CreateWebSite(string) Website
	PrintWebSite()
}

func main() {
	log.Println(" Starting ..")
	// Get config
	config := getConfig()

	// Create a website object
	website := web.CreateWebSite(config.BaseURL.String())

	// Create crawler object
	crawler := crawler.NewCrawler(config, website)

	log.Println("====== Starting crawling =====")
	crawler.Crawl(config.BaseURL.String(), config.Depth)
	log.Println("====== Done  crawling =====")

	// Print website
	website.PrintBasicSiteMap()
	website.PrintSiteGraph()
}

func getConfig() crawler.Config {
	cfg := crawler.Config{}
	cfg.Concurrency = 2
	cfg.Depth = 2
	cfg.BaseURL, _ = url.Parse("https://monzo.com/")
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
