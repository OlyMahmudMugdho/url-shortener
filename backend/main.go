package main

import (
	"github.com/OlyMahmudMugdho/url-shortener/cmd/api"
)

func main() {
	server := api.NewApiServer(":8080")
	server.Run()
}
