package scrawler

import (
	"fmt"
	"log"
)

// DS
type Website interface {
	AddWebpage(Webpage)
	// CheckIfWebpageExists(wpage *Webpage)
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
	log.Println("Printing Sitemap")
	fmt.Println(" ====  Sitemap ==== ")
	log.Println(len(w.webpages))
	for _, page := range w.webpages {
		fmt.Println(page.String())
	}
	log.Println(" ==== Sitemap Done  ==== ")
}

func (w *website) PrintSiteGraph() {
	log.Println("Printing SiteGraph")
	fmt.Println(" ====  Sitemap ==== ")
	log.Println(len(w.webpages))
	for _, page := range w.webpages {
		fmt.Println(page.String())
		for _, reference := range page.references {
			fmt.Println("==> " + reference.String())
		}
	}
	fmt.Println(" ==== SiteGraph Done  ==== ")
	log.Println("SiteGraph done")
}

// function
func CreateWebSite(sitename string) *website {
	wsite := website{sitename, nil}
	return &wsite
}
