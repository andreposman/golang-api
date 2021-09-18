package main

import (
	"github.com/andreposman/golang-api/cmd/api"
)


func main() {
	/*
	TODO:
		- add tests
		- validate requirements
		- add logger
		- improve req validations
		- add swagger
		- better data manipulation
		- change lat/lng to strings
		- docker/docker-compose
		- review solid/12factors/clean code principles

	TODO improvements:
		- host cloud
		- add db
		- monitor/observability tools

	*/
	api.InitAPI()
}
