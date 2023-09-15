package config

import (
	"fmt"
	"testing"
)

func TestInitConfigFromFile(t *testing.T) {
	InitConfigFromFile("../config.ini")
	fmt.Println(DataBaseType)
	fmt.Println(DataBaseName)
	fmt.Println(Username)
}

func TestGetDsn(t *testing.T) {
	InitConfigFromFile("../config.ini")
	fmt.Println(GetDsn())
}
