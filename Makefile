build-image:
	docker build . -t rest-api

up:
	docker build . -t rest-api
	docker-compose -f docker-compose.yml -p services up --build --detach

down:
	docker-compose -f docker-compose.yml -p services down --remove-orphans

docker-clean:
	docker system prune --volumes --force

swagger-docs:
	swag init -d internal/handler -g main.go -o api/swagger --parseDependency
	
test:
	go test ./...

file_tree:
	tree -I vendor