
APOLLOBIN=ApolloStats

$(APOLLOBIN):
	go build -o $(APOLLOBIN)
	cd src/ && rice append --exec ../$(APOLLOBIN)

build: $(APOLLOBIN)

clean:
	rm $(APOLLOBIN)

