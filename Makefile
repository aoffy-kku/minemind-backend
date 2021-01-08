dev:
	swag init
	go run main.go
test:
	go run cmd/test/main.go
migrate:
	go run cmd/migrate/main.go
build:
	docker build -f Dockerfile -t asia.gcr.io/kube101-292215/minemind-backend .
	docker push asia.gcr.io/kube101-292215/minemind-backend