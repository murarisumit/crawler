Web-crawler

Depedency : 
* goquery - a little like that j-thing, only in Go: `go get github.com/PuerkitoBio/goquery`
* rlog: - A simple Golang logger : `github.com/romana/rlog`

### Run the application
1. Clone the repo   : `git clone github.com/murarisumit/crawler.git`
2. Get dependencies : `make get-deps`
3. Build and Run    : `make run`

### Output
File is created with name: `sitemap.txt` and `sitegraph.txt`


### Logging

DEBUG/INFO logs can be enabled by setting `RLOG_LOG_LEVEL` env variable to "DEBUG", "INFO", "WARN", "ERROR", "CRITICAL" or "NONE"

More refeence about logging: https://github.com/romana/rlog#controlling-rlog-through-environment-or-config-file-variables

