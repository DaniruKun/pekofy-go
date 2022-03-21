build:
	go build -o pekofy .

run:
	go run main.go

test:
	go test

all: test build