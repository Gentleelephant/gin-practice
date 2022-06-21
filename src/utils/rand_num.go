package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func RandNum(n int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	if n <= 0 || n > 9 {
		n = 6
	}
	format := "%0" + fmt.Sprintf("%d", n) + "v"
	vcode := fmt.Sprintf(format, rnd.Int31n(1000000))
	return vcode
}
