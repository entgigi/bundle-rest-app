build:
	go build -o bin/bundle-rest-app main.go

run: build
	go run main.go	