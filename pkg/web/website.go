package web

import (
	// "fmt"
	"io"
	"log"
	"os"
	"strings"
)

// A webpage is url with refrences to other webpages
type Webpage struct {
	URL        string
	References []Webpage
}

// A website is a list of webpages
type website struct {
	name     string
	webpages []Webpage
}

// method
func (w *website) AddWebpage(wpage Webpage) {
	w.webpages = append(w.webpages, wpage)
}

func (w *website) PrintBasicSiteMap() {
	log.Println("==== Printing Sitemap ====")

	log.Println("Number of webpages: " + string(len(w.webpages)))
	log.Println("Output create in file sitemap.txt")

	file, err := os.Create("sitemap.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	for _, page := range w.webpages {
		io.Copy(file, strings.NewReader(page.URL+"\n"))
		// io.Copy(file, strings.NewReader(page.String()))
	}

	log.Println("==== Sitemap Done  ==== ")
}

func (w *website) PrintSiteGraph() {
	log.Println("====  Printing SiteGraph ====")
	log.Println("Output create in file \"sitegraph.txt\"")

	file, err := os.Create("sitegraph.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	for _, page := range w.webpages {
		// fmt.Fprintln(file, page.String())
		io.Copy(file, strings.NewReader(page.URL+"\n"))
		for _, reference := range page.References {
			io.Copy(file, strings.NewReader("-> "+reference.URL+"\n"))
			// fmt.Fprintln(file, "-> "+reference.String())
		}
	}
	log.Println("==== SiteGraph done ==== ")
}

func CreateWebSite(siteurl string) *website {
	wsite := website{siteurl, nil}
	return &wsite
}
