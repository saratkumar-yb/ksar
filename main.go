package main

import (
	apiv1 "ksar/api"
	"os"
)

func main() {
	a := apiv1.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8080")
}
