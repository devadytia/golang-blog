package migrations

import (
	"blog/database"
)

func Migrate() {
	database.ConnectDB()
	Migration()
}
