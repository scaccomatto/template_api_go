build-image:
	docker build . -t template-api

up:
	docker build . -t template-api
	docker-compose -f docker-compose.yml -p templateapi up

down:
	docker-compose -f docker-compose.yml -p templateapi down --remove-orphans

docker-clean:
	docker system prune --volumes --force

swagger-docs:
	swag init -d internal/handler -g main.go -o api/swagger --parseDependency
	
test:
	go test ./...

file_tree:
	tree -I vendor