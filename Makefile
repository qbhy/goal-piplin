DOCKER_TAG=goal
CONSOLE=go run bootstrap/console/main.go

run:
	go run bootstrap/app/main.go run

build:
	go build -o ./piplin -v ./

test:
	go test -json ./tests

pack:
	docker build -t $(DOCKER_TAG) .

migrate:
	$(CONSOLE) migrate

migrate-rollback:
	$(CONSOLE) migrate:rollback

migrate-refresh:
	$(CONSOLE) migrate:refresh

migrate-reset:
	$(CONSOLE) migrate:reset

migrate-status:
	$(CONSOLE) migrate:status

make-migration:
	$(CONSOLE) make:migration $(NAME)

install: migrate
	$(CONSOLE) init

update:
	docker compose pull views server
	docker compose up -d views server
