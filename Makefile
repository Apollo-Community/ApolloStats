
APOLLOBIN=ApolloStats
LINUX32=env GOOS=linux GOARCH=386
LINUX64=env GOOS=linux GOARCH=amd64

build: build64

build32: generate
	$(LINUX32) go build -tags "embed" -o "$(APOLLOBIN)32"

build64: generate
	$(LINUX64) go build -tags "embed" -o "$(APOLLOBIN)64"

generate:
	cd src/assettemplates/  && yaber --pkg "assettemplates" -strip "../../"  ../../templates/
	cd src/assetstatic && yaber --pkg "assetstatic" -strip "../../" ../../static/

test_run:
	go run main.go run

test_run_embed:
	go run -tags "embed" main.go run

clean:
	rm "$(APOLLOBIN)32"
	rm "$(APOLLOBIN)64"

