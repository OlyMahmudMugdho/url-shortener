package main

import (
	"fmt"

	"github.com/OlyMahmudMugdho/gotenv/gotenv"
	"github.com/OlyMahmudMugdho/url-shortener/utils"
)

func main() {
	gotenv.Load()
	if utils.ConnectToDatabase() {
		fmt.Println("connected to the database")
	}
}
