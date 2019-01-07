package crawler

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/murarisumit/crawler/pkg/web"
)

type Website interface {
	AddWebpage(web.Webpage)
}

type crawler struct {
	Config // has a config object to maintain a buffer to urls to be parsed
	// Want to make time-limited calls to server, hence mainting a bufferd queue
	// It gives me flexibility from thread prevention and things.
	//queue    chan Webpage       // not implemented
	website Website         // crawler has a reference to website obj
	crawled map[string]bool // to check page is crawled or not
}

// Crawl uses fetcher to recursively crawl
func (c *crawler) Crawl(inptURL string, depth int) {
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
		// parsedURL, _ := url.Parse(inptURL) //parsedURL
		currentPage := web.Webpage{inptURL, nil}
		urls, _ := parseWebPage(inptURL)
		for _, u := range urls {
			parsedRef, _ := url.Parse(u)
			// too many webpage object,
			// not sure if creating webpage obj for other domain too makes sense
			wpage := web.Webpage{parsedRef.String(), nil}
			currentPage.References = append(currentPage.References, wpage)
		}
		log.Println("Added all references of : " + inptURL)
		c.crawled[inptURL] = true
		c.website.AddWebpage(currentPage)

		for _, u := range urls {
			c.Crawl(u, depth-1)
		}
	}
	return
}

// Check if request comes is excluded list
func (c crawler) isExcluded(inptURL string) bool {
	href, _ := url.Parse(inptURL)

	// Gate 1: Hostname doesn't have suffix "monzo.com" return
	if !strings.HasSuffix(href.Hostname(), c.BaseURL.Hostname()) {
		log.Println(c.BaseURL.Hostname() + " : doesn't match with : " + href.Hostname())
		return true
	}

	//Gate 2: Path should contain one of those string.
	path := href.Path
	for _, v := range c.ExcludedPath {
		if strings.HasPrefix(path, v) {
			log.Println(href.Path + " : prefix is in excluded list")
			return true
		}
	}

	//Gate 3: Should not be part of subdomain
	domain := href.Hostname()
	for _, v := range c.ExcludedSubdomain {
		if strings.Contains(domain, v) {
			log.Println(domain + ": is part of excluded list")
			return true
		}
	}
	return false
}

// Fetches a webpage and all references it has.
// This function doesn't care of it's excluded or not,
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

func NewCrawler(config Config, website Website) *crawler {
	c := &crawler{}
	c.Config = config
	c.website = website
	c.crawled = make(map[string]bool)
	return c
}
