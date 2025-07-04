# Golang rest api example

This is just and example of a simple API that I can use as template.  
This is a "layered architecture" application. Here a quick description:
* cmd: it contains the main file that points to the application staring point.
* api: it contains swagger docs.
* internal: it contains all code that I want to keep it private: business logic, services, etc...

```shell
├── api
│        └── swagger
│            ├── docs.go
│            ├── readme.md
│            ├── swagger.json
│            └── swagger.yaml
├── cmd
│        └── main.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│        ├── app
│        │        └── app.go
│        ├── apperr
│        │        └── apperr.go
│        ├── appmiddleware
│        │        └── middleware.go
│        ├── conf
│        │        ├── conf.go
│        │        └── config.yaml
│        ├── handler
│        │        ├── main.go
│        │        └── user.go
│        ├── logger
│        │        └── logger.go
│        ├── model
│        │        └── user.go
│        ├── repository
│        │        ├── gormdb.go
│        │        ├── schema.sql
│        │        ├── user.go
│        │        └── user_test.go
│        └── service
│            ├── users.go
│            └── user_test.go
├── Makefile
└── README.md

```




## How to

In the Makefile there are a few shortcuts to build and running the server using docker and docker compose  
You can find Swagger documentation in api folder 


