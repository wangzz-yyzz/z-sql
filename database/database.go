package database

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"z-sql/config"
)

type Connection struct {
	DataBase *gorm.DB
	SqlDb    *sql.DB
}

// InitDataBase
// open database and set max idle and max open connections
func InitDataBase() *Connection {
	// init error
	var err error

	// open database
	DataBase, err := gorm.Open(mysql.Open(config.GetDsn()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// open sql database
	SqlDb, err := DataBase.DB()
	if err != nil {
		panic(err)
	}

	// set max idle and max open connections
	SqlDb.SetMaxIdleConns(config.MaxIdleCons)
	SqlDb.SetMaxOpenConns(config.MaxOpenCons)

	return &Connection{
		DataBase: DataBase,
		SqlDb:    SqlDb,
	}
}

func InitDataBaseFromConfig(myConfig config.Config) *Connection {
	// init error
	var err error

	// open database
	DataBase, err := gorm.Open(mysql.Open(config.GetDsnFromConfig(myConfig)), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// open sql database
	SqlDb, err := DataBase.DB()
	if err != nil {
		panic(err)
	}

	// set max idle and max open connections
	SqlDb.SetMaxIdleConns(config.MaxIdleCons)
	SqlDb.SetMaxOpenConns(config.MaxOpenCons)

	return &Connection{
		DataBase: DataBase,
		SqlDb:    SqlDb,
	}
}
