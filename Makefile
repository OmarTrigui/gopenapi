install:
	go install -v


fmt:
	go fmt


gen: validate
	rm -rf ./pkg/swagger/server/*
	swagger generate server \
		-t ./pkg/swagger/server/ \
		-f ./pkg/swagger/swagger.yml \
		--exclude-main \
		-A gopenapi

run:
	go run internal/main.go

validate:
	swagger validate ./pkg/swagger/swagger.yml

build:
	go build -o bin/gopenapi internal/main.go

doc:
	docker run --rm -i yousan/swagger-yaml-to-html < pkg/swagger/swagger.yml > doc/index.html

.PHONY: install fmt gen validate
