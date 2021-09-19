package main

import (
	"github.com/andreposman/golang-api/internal/api"
	"github.com/andreposman/golang-api/internal/helpers/banner"
)


func main() {
	/*
	TODO:
		- add tests
		- validate requirements
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

	banner.GreetingBanner()
	api.InitAPI()
}
