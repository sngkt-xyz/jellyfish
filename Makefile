.PHONY: build
build:
	go build -o ./tmp/main .

.PHONY: start-dev
start-dev:
	air -c .air.toml

.PHONY: start
start: build
	chmod +x ./tmp/main
	./tmp/main
