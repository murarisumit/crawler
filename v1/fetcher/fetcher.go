package fetcher

import (
	"io"
	"log"
	"net/http"
)

func Hello() string {
	return "sumit"
}

func FetchWebPage(url string) (io.ReadCloser, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
			return "", res.Status
		}
		return res.Body, nil
	}
}
