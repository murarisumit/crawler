package scrawler

import (
	// "fmt"
	"io"
	"log"
	"os"
	"strings"
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
	defer file.Close()
	if err != nil {
		panic(err)
	}
	for _, page := range w.webpages {
		io.Copy(file, strings.NewReader(page.String()+"\n"))
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
		io.Copy(file, strings.NewReader(page.String()+"\n"))
		for _, reference := range page.references {
			io.Copy(file, strings.NewReader("-> "+reference.String()+"\n"))
			// fmt.Fprintln(file, "-> "+reference.String())
		}
	}
	log.Println("==== SiteGraph done ==== ")
}

// function
func CreateWebSite(sitename string) *website {
	wsite := website{sitename, nil}
	return &wsite
}
