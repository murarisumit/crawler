package scrawler

import "net/url"

type Config struct {
	//Main url we are parsing
	BaseURL *url.URL
	// number of concurrent calls
	Concurrency int
	// time before making next call
	waitTime int
	// depth to visit
	Depth int
	// excluded endpoint
	excluded []string
}
