build:
	go build

clean:
	rm crawler

get-deps:
	go get github.com/PuerkitoBio/goquery

run: | build
	./crawler


