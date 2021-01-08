dev:
	swag init
	go run main.go
test:
	go run cmd/test/main.go
migrate:
	go run cmd/migrate/main.go
build:
	docker build -f Dockerfile -t minemind-backend .