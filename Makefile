run-cli: vendor
	go run ./cmd/cli/*.go -config ./data/config.yml

run-http: vendor
	go run ./cmd/server/*.go -config ./data/config.yml

build: vendor
	CGO_ENABLED=1 go build -mod=vendor -v -o bin/server ./cmd/server

clean-db:
	rm ./data/goticker.db

vendor:
	go mod vendor

watch-client:
	cd cmd/server/server-app && npm run watch

watch:
	modd

.PHONY: vendor