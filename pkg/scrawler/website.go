package scrawler

import (
	"fmt"
	"log"
	"os"
)

// DS
type Website interface {
	AddWebpage(Webpage)
	// CheckIfWebpageExists(wpage Webpage)
	PrintBasicSiteMap()
	PrintSiteGraph()
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
	if err != nil {
		panic(err)
	}
	for _, page := range w.webpages {
		fmt.Fprintln(file, page.String())
	}

	log.Println("==== Sitemap Done  ==== ")
}

func (w *website) PrintSiteGraph() {
	log.Println("====  Printing SiteGraph ====")
	log.Println("Output create in file \"sitegraph.txt\"")

	file, err := os.Create("sitegraph.txt")
	if err != nil {
		panic(err)
	}
	for _, page := range w.webpages {
		fmt.Fprintln(file, page.String())
		for _, reference := range page.references {
			fmt.Fprintln(file, "-> "+reference.String())
		}
	}
	log.Println("==== SiteGraph done ==== ")
}

// function
func CreateWebSite(sitename string) *website {
	wsite := website{sitename, nil}
	return &wsite
}
