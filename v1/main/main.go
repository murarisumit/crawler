package main

import (
	// "crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	// "github.com/murarisumit/crawler/pkg/fetcher"
	// "github.com/murarisumit/crawler/pkg/parser"
	"github.com/PuerkitoBio/goquery"
)

type SiteGraph struct {
	visited map[url.URL]bool
	sitemap map[url.URL][]url.URL
}

// type node struct {
// 	url.URL
// 	isCrawled bool
// }

var ignorePath = []string{
	"/cdn-cgi",
	"/legal",
	"/static",
}

var ignoreSubdomain = []string{
	"www.monzo.com",
}

var sitegraph SiteGraph

func main() {
	raw_endpoint := "https://monzo.com"
	endpoint, err := url.Parse(raw_endpoint)
	if err != nil {
		log.Fatal("URL parsing error")
	}
	sitegraph = SiteGraph{}
	sitegraph.visited = make(map[url.URL]bool)
	sitegraph.sitemap = make(map[url.URL][]url.URL)
	sitegraph.visited[*endpoint] = false

	crawl(endpoint)
	printSiteGraph()

}

func crawl(endpoint *url.URL) {
	res, _ := http.Get(endpoint.String())
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("body a").Each(func(i int, s *goquery.Selection) {
		raw_href, exists := s.Attr("href")
		if exists {
			href, _ := url.Parse(raw_href)
			if strings.HasPrefix(raw_href, "/") {
				href = endpoint.ResolveReference(href)
			}
			// Gate 1: should be of same hostname
			if href.Hostname() == endpoint.Hostname() {
				//Gate 2: Should not be visited
				if _, visted := sitegraph.visited[*href]; !visted {
					//Gate 3: Path should contain one of those string.
					if !excluded(href) {
						fmt.Println("New link : " + href.String())
						if _, present := sitegraph.visited[*href]; present {
							sitegraph.sitemap[*endpoint] = append(sitegraph.sitemap[*endpoint], *href)
						} else {
							sitegraph.visited[*href] = false
							sitegraph.sitemap[*endpoint] = append(sitegraph.sitemap[*endpoint], *href)
						}
						crawl(href)
						sitegraph.visited[*href] = true
					}
				}
			}
		}
	})
}

func printSiteGraph() {
	fmt.Println("========== Sitemap =================")
	for k, v := range sitegraph.visited {
		fmt.Printf("%s : %t \n", k.String(), v)
		for i2, v2 := range sitegraph.sitemap[k] {
			fmt.Printf("=== %d : %s \n", i2, v2.String())
		}
	}
	fmt.Println("========== Sitemap =================")
}

// Check if request path contains excluded url
func excluded(href *url.URL) bool {

	path := href.Path
	for _, v := range ignorePath {
		if strings.HasPrefix(path, v) {
			return true
		}
	}

	domain := href.Hostname()
	for _, v := range ignoreSubdomain {
		if strings.Contains(domain, v) {
			return true
		}
	}

	return false
}
