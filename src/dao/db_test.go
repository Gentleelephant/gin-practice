package dao

import (
	"gin-practice/src/config"
	"testing"
)

func TestCreateDatabase(t *testing.T) {
	err := config.LoadConfig(config.ConfigPath)
	if err != nil {
		return
	}
	config.DB = InitDB()
	if err != nil {
		return
	}
}
