build:
	go build

clean:
	rm crawler

get-deps:
	go get github.com/PuerkitoBio/goquery
	go get github.com/romana/rlog

run: | build
	./crawler


