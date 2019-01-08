package main

import (
	"net/url"

	"github.com/murarisumit/crawler/pkg/crawler"
	"github.com/murarisumit/crawler/pkg/web"
	log "github.com/romana/rlog"
)

type Crawler interface {
	Crawl(string, int)
}

type Website interface {
	CreateWebSite(string) Website
	PrintWebSite()
}

func getConfig() crawler.Config {
	cfg := crawler.Config{}
	cfg.Concurrency = 1
	cfg.Depth = 1
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

func main() {
	// Get config
	config := getConfig()

	// Create a website object
	website := web.CreateWebSite(config.BaseURL.String())

	// Create crawler object
	crawler := crawler.NewCrawler(config, website)

	log.Info("====== Starting crawling =====")
	crawler.Crawl(config.BaseURL.String(), config.Depth)
	log.Info("====== Done  crawling =====")

	// Print website
	website.PrintBasicSiteMap()
	website.PrintSiteGraph()
}
