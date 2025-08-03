APP_NAME=lu_api
DOCKER_IMAGE=lu_api_img
DOCKER_CONTAINER=lu_API
DOCKER_COMPOSE=docker-compose
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
	$(DOCKER_COMPOSE) up --build

down:
	$(DOCKER_COMPOSE) down

