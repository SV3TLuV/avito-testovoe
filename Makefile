DOCKER_IMAGE_NAME=avito-zadanie-6105
DOCKER_PORT=8080

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

docker-run:
	docker run -d -p $(DOCKER_PORT):8080 $(DOCKER_IMAGE_NAME)

docker-stop:
	docker stop $(shell docker ps -q --filter "ancestor=$(DOCKER_IMAGE_NAME)")
	docker rm $(shell docker ps -a -q --filter "ancestor=$(DOCKER_IMAGE_NAME)")

migrate:
	go run ./src/cmd/migrate/main.go

run:
	go run ./src/cmd/api/main.go

test:
	venom run venom-tests.yml

test-without-logs:
	make test clean

clean:
	rm -f venom.*.log venom.log
