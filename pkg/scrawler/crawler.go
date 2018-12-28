package scrawler

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var ignorePath = []string{
	"/cdn-cgi",
	"/legal",
	"/static",
}

var ignoreSubdomain = []string{
	"www.monzo.com",
}

type Crawler interface {
	Print()
	Crawl(string, int)
}

type crawler struct {
	// to maintain a buffer to urls to be parsed
	// Want to make time-limited calls to server, hence mainting a bufferd queue
	// It gives me flexibility from thread prevention and things.
	//queue    chan Webpage       // not implemented
	refStore map[string]Webpage // to store the reference of webpages
	crawled  map[string]bool    // to check page is crawled or not

	config  Config  // has a config object
	website Website // crawler has a interface to website object
}

// Public method
func (c crawler) Print() {
	log.Println("Printing crawler")
	c.website.PrintBasicSiteMap()
	c.website.PrintSiteGraph()
}

// Crawl uses fetcher to recursively crawl
func (c crawler) Crawl(inptURL string, depth int) {
	// To till certain depth
	if depth <= 0 {
		return
	}

	if c.crawled[inptURL] {
		log.Println("Already crawled : " + inptURL)
	}

	// If page is not crawled till now and should not be in excluded url
	if !c.crawled[inptURL] && !c.isExcluded(inptURL) {
		log.Println("Starting for : " + inptURL)
		parsedURL, _ := url.Parse(inptURL) //parsedURL
		currentPage := Webpage{*parsedURL, nil}
		urls, _ := parseWebPage(inptURL)
		// log.Println("Got urls for : " + inptURL)
		// log.Println(urls)
		for _, u := range urls {
			parsedRef, _ := url.Parse(u)
			wpage := Webpage{*parsedRef, nil}
			// log.Println("webpage object created")
			c.refStore[inptURL] = wpage
			currentPage.addReference(wpage)
			// log.Println("Added reference to " + currentPage.String())
			// c.refStore[inptURL].addReference(wpage)
		}
		log.Println("Added all references to : " + currentPage.String())
		c.crawled[inptURL] = true
		c.website.AddWebpage(currentPage)

		for _, u := range urls {
			c.Crawl(u, depth-1)
		}
	}
	return
}

// Check if request path contains excluded url
func (c crawler) isExcluded(inptURL string) bool {
	href, _ := url.Parse(inptURL)
	// Gate 1: Hostname doesn't have suffix "monzo.com" return
	// if !strings.HasSuffix(href.Hostname(), "monzo.com") {
	if !strings.HasSuffix(href.Hostname(), c.config.BaseURL.Hostname()) {
		// log.Println(c.config.BaseURL.Hostname())
		log.Println("HostName doesn't end with monzo: " + inptURL)
		return true
	}

	//Gate 2: Path should contain one of those string.
	path := href.Path
	for _, v := range ignorePath {
		if strings.HasPrefix(path, v) {
			return true
		}
	}

	//Gate 3: Should not be part of subdomain
	domain := href.Hostname()
	for _, v := range ignoreSubdomain {
		if strings.Contains(domain, v) {
			return true
		}
	}
	return false
}

//Fetches a webpage and all references it has.
//This function doesn't care of it's excluded or not, it just pulls all the links in page
func parseWebPage(inptURL string) (referenced []string, err error) {
	output := []string{}
	parsedURL, _ := url.Parse(inptURL)

	res, _ := http.Get(inptURL)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalln("Error while fetching body for : " + inptURL)
		log.Fatal(err)
	}

	doc.Find("body a").Each(func(i int, s *goquery.Selection) {
		raw_href, exists := s.Attr("href")
		if exists {
			href, _ := url.Parse(raw_href)
			if strings.HasPrefix(raw_href, "/") ||
				strings.HasPrefix(raw_href, "#") ||
				strings.HasPrefix(raw_href, ".") {
				href = parsedURL.ResolveReference(href)
			}
			output = append(output, href.String())
		}
	})
	return output, nil
}

// We are forced to call the constructor to get an instance of candidate
func NewCrawler(config Config) Crawler {
	c := crawler{}
	c.config = config
	c.website = CreateWebSite(config.BaseURL.String())
	// c.queue = make(chan Webpage, config.Concurrency) // Set buffer length for our queue
	c.refStore = make(map[string]Webpage)
	c.crawled = make(map[string]bool)
	return c
}
