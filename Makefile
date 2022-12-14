migrate:
	migrate -path ./internal/migrations -database "postgres://postgres:qwerty@0.0.0.0:5433/simple?sslmode=disable" up
cmigrate:
	migrate create -ext sql -dir ./internal/migrations -seq $(name)
fmigrate:
	migrate -path ./internal/migrations -database "postgres://postgres:qwerty@0.0.0.0:5433/simple?sslmode=disable" force $(id)
dmigrate:
	migrate -path ./internal/migrations -database "postgres://postgres:qwerty@0.0.0.0:5433/simple?sslmode=disable" down 1
swag:
	swag init -g ./cmd/app/main.go
build:
	docker-compose build simple-app
run:
	docker-compose up -d simple-app
build_run: build run