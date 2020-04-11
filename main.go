package main

import (
	"flag"
	"fmt"
	"my-gin-boilerplate/config"
	"my-gin-boilerplate/db"
	"my-gin-boilerplate/routes"
	"os"
)

func main() {
	environment := flag.String("e", "development", "define environment mode development or production. (default development)")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)

	db.Init()
	defer db.CloseDB()

	routes.Init()
}
