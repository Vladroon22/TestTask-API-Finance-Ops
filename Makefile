.PHONY:

run:
	docker-compose up --build

app:
	goose up
	go build -o ./app cmd/main.go
	./app
