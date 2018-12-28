package scrawler

import (
	"net/url"
)

// A webpage is url with refrences to other webpages
type Webpage struct {
	url.URL
	references []Webpage
}

func (w *Webpage) create(addr url.URL) Webpage {
	wpage := Webpage{addr, nil}
	return wpage
}

func (w *Webpage) addReference(ref Webpage) {
	w.references = append(w.references, ref)
}
