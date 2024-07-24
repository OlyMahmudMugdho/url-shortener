package main

import (
	"github.com/OlyMahmudMugdho/gotenv/gotenv"
	"github.com/OlyMahmudMugdho/url-shortener/utils"
)

func main() {
	gotenv.Load()
	utils.ConnectToDatabase()
}
