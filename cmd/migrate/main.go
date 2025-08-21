package main

import (
	"blog/database"
	"blog/database/migrations"
)

func main() {
	database.ConnectDB()
	migrations.Migration()
}
