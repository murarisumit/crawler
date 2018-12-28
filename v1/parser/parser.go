package parser

import (
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"
)

// func ParseWebPage(page string) (urls []string, err error) {
func ParseWebPage(page io.ReadCloser) {
	fmt.Println(page)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
	})
}
