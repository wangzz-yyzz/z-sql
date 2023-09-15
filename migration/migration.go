package migration

import (
	"log"
	"z-sql/database"
)

// CreateTableIfNotExist
// create table if not exist
func CreateTableIfNotExist(connection *database.Connection, model any) error {
	log.Println("creating table...")
	migrator := connection.DataBase.Migrator()
	if migrator.HasTable(model) {
		log.Println("exist table")
		return nil
	} else {
		return connection.DataBase.AutoMigrate(model)
	}
}

// ReCreateTable
// drop table if exist and create table
func ReCreateTable(connection *database.Connection, model any) error {
	migrator := connection.DataBase.Migrator()
	if migrator.HasTable(model) {
		log.Println("exist table, dropping...")
		err := migrator.DropTable(model)
		if err != nil {
			return err
		}
	}

	return CreateTableIfNotExist(connection, model)
}
