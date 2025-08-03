APP_NAME=lu_api
DOCKER_IMAGE=lu_api_img
DOCKER_CONTAINER=lu_API
PORT=8000

build:
	go build -o $(APP_NAME) ./cmd/api

run:
	go run ./cmd/api

docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	docker run -p $(PORT):$(PORT) $(DOCKER_IMAGE)

up:
	docker-compose up --build

down:
	docker-compose down

