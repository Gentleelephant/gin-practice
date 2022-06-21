package mail

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"os"
	"sync"
	"time"
)

func main() {
	ch := make(chan *email.Email, 10)
	p, err := email.NewPool(
		"smtp.qq.com:25",
		5,
		smtp.PlainAuth("", "1665400978@qq.com", "srwrbfpblphtbcac", "smtp.qq.com"),
	)
	if err != nil {
		return
	}

	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func() {
			defer wg.Done()
			for e := range ch {
				err := p.Send(e, 10*time.Second)
				if err != nil {
					fmt.Fprintf(os.Stderr, "email:%v sent error:%v\n", e, err)
				}
			}
		}()
	}

	for i := 0; i < 10; i++ {
		e := email.NewEmail()
		e.From = "dj <XXX@qq.com>"
		e.To = []string{"XXX@qq.com"}
		e.Subject = "Awesome web"
		e.Text = []byte(fmt.Sprintf("Awesome Web %d", i+1))
		ch <- e
	}

	close(ch)
	wg.Wait()
}
