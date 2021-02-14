dev:
	swag init
	go run main.go
test:
	go run cmd/test/main.go
migrate:
	go run cmd/migrate/main.go
build:
	docker build -f Dockerfile -t minemind-backend .
run:
	docker run --name minemind-backend -p 1321:1321 -d minemind-backend
deploy:
	docker build -f Dockerfile -t minemind-backend .
	make restart
restart:
	docker stop minemind-backend
	docker rm minemind-backend
	docker run --name minemind-backend -p 1321:1321 -d minemind-backend