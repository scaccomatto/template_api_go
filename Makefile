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
	swag init -d internal/app/httpserver -g httpserver.go -o api/swagger --parseVendor --parseDependency --md ./docs/
	
test:
	go test ./...