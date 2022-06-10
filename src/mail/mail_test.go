package mail

import (
	"log"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

func TestMail(t *testing.T) {

	e := email.NewEmail()

	// 设置发送方邮箱
	e.From = "1665400978@qq.com"

	// 设置接收方邮箱
	e.To = []string{"1132960613@qq.com"}

	// 设置抄送
	//e.Cc = []string{"XXX@qq.com",XXX@qq.com}
	// 设置秘密抄送
	//e.Bcc = []string{"XXX@qq.com"}

	//设置文件发送的内容
	//e.HTML = []byte(`
	//<h1><a href="http://www.topgoer.com/">go语言中文网站</a></h1>
	//`)
	//这块是设置附件
	//e.AttachFile("./test.txt")

	// 设置主题
	e.Subject = "这是主题"

	// 设置内容
	e.Text = []byte("www.topgoer.com是个不错的go语言中文文档")

	// 设置服务器配置并发送
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "1665400978@qq.com", "srwrbfpblphtbcac", "smtp.qq.com"))

	if err != nil {
		log.Fatal(err)
	}
}
