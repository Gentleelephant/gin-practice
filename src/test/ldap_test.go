package test

import (
	"gin-practice/src/auth"
	"gin-practice/src/config"
	"testing"
)

func TestLdap(t *testing.T) {

	err := config.LoadConfig(config.ConfigPath)
	if err != nil {
		return
	}
	provider := auth.GetProvider()
	authentication := provider.Authentication("zhang", "zhang")
	if authentication == false {
		t.Error("authentication failed")
	}
}
