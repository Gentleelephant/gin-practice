package initialize

import (
	"gin-practice/src/global"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func InitEmailPool() {
	config := global.GlobalConfig.EmailConfig
	pool, err := email.NewPool(config.Address, config.Count, smtp.PlainAuth("", config.Username, config.Password, config.Host))
	if err != nil {
		//
	}
	global.Pool = pool

	emails := make(chan *email.Email, 64)
	global.EmailLists = emails
}
