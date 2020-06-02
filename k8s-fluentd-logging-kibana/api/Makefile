.PHONY: build
build:
	docker build -t rails-api .

.PHONY: run
run: build
	docker run -p 8080:8080 rails-api
