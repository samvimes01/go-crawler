PHONY: run, build, buildrun, test

test: 
	go test ./...	

run: 
	go run . "https://www.wagslane.dev" 6 20

build: 
	rm -f crawler && go build -o crawler .

buildrun: build
	./crawler "https://www.wagslane.dev" 6 20