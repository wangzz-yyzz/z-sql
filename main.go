package main

import (
	"database/sql"
	"z-sql/config"
	"z-sql/database"
	"z-sql/migration"
)

type User struct {
	Name string
	Age  int
}

func main() {

	myConfig := config.InitConfigFromFile("config.ini")

	connection := database.InitDataBaseFromConfig(myConfig)

	defer func(SqlDb *sql.DB) {
		err := SqlDb.Close()
		if err != nil {
			panic(err)
		}
	}(connection.SqlDb)

	err := migration.CreateTableIfNotExist(connection, &User{})
	if err != nil {
		panic(err)
	}

	err = migration.ReCreateTable(connection, &User{})
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "hangman",
		Age:  18,
	}

	// insert
	connection.DataBase.Select("*").Create(&user)

	// retrieve
	var users []User
	connection.DataBase.Select("*").Find(&users)
	for _, user := range users {
		println(user.Name, user.Age)
	}

	// update
	connection.DataBase.Model(&user).Where("name = ?", "hangman").Update("age", 20)
	connection.DataBase.Select("*").Find(&users)
	for _, user := range users {
		println(user.Name, user.Age)
	}

	// delete
	connection.DataBase.Where("name = ?", "hangman").Delete(&User{})
	connection.DataBase.Select("*").Find(&users)
	for _, user := range users {
		println(user.Name, user.Age)
	}
}
