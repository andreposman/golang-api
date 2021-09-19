package main

import (
	"github.com/andreposman/golang-api/internal/api"
	"github.com/andreposman/golang-api/internal/helpers/banner"
)


func main() {
	banner.GreetingBanner()

	api.InitAPI()
}
