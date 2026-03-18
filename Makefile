run:
	go run cmd/mock-cbr/main.go

build:
	go build -o app cmd/mock-cbr/main.go

swagger:
	swag init -g cmd/mock-cbr/main.go

redis-up:
	docker-compose up -d redis

redis-run: redis-up
	$(MAKE) run