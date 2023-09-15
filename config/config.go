package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
)

var (
	DataBaseType string
	DataBaseName string
	Username     string
	Password     string
	Host         string
	Port         string
	MaxIdleCons  int
	MaxOpenCons  int
)

type Config struct {
	DataBaseType string
	DataBaseName string
	Username     string
	Password     string
	Host         string
	Port         string
	MaxIdleCons  int
	MaxOpenCons  int
}

func InitConfigFromFile(path string) Config {
	// load config file
	config, err := ini.Load(path)
	if err != nil {
		log.Println("load config file error")
		panic(err)
	}

	// get config from config file
	DataBaseType = config.Section("").Key("DataBaseType").MustString("mysql")
	DataBaseName = config.Section("").Key("DataBaseName").MustString("test")
	Username = config.Section("").Key("Username").MustString("root")
	Password = config.Section("").Key("Password").MustString("password")
	Host = config.Section("").Key("Host").MustString("127.0.0.1")
	Port = config.Section("").Key("Port").MustString("3306")
	MaxIdleCons = config.Section("").Key("MaxIdleCons").MustInt(10)
	MaxOpenCons = config.Section("").Key("MaxOpenCons").MustInt(100)

	log.Println("config init success")

	return Config{
		DataBaseType: DataBaseType,
		DataBaseName: DataBaseName,
		Username:     Username,
		Password:     Password,
		Host:         Host,
		Port:         Port,
		MaxIdleCons:  MaxIdleCons,
		MaxOpenCons:  MaxOpenCons,
	}
}

// GetDsn
// get dsn string from config
func GetDsn() string {
	if DataBaseType == "mysql" {
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", Username, Password, Host, Port, DataBaseName)
	} else if DataBaseType == "sqlite" {
		return ""
	}

	log.Println("database type not supported")
	return ""
}

func GetDsnFromConfig(config Config) string {
	if config.DataBaseType == "mysql" {
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			config.Username, config.Password, config.Host, config.Port, config.DataBaseName)
	} else if config.DataBaseType == "sqlite" {
		return ""
	}

	log.Println("database type not supported")
	return ""
}
