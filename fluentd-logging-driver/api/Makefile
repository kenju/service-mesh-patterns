.PHONY: build
build:
	docker build -t sinatra-api .

.PHONY: run
run: build
	docker run -p 4567:4567 sinatra-api
