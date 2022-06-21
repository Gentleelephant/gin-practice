package v1

import (
	"encoding/base64"
	"errors"
	"gin-practice/src/global"
	"gin-practice/src/utils"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"io/ioutil"
	"os"
	"strings"
)

type RespCaptcha struct {
	CaptchaId string `json:"captchaId"`
	Base64    string `json:"base64"`
}

type ReqCaptcha struct {
	CaptchaId string `json:"captchaId"`
	Captcha   string `json:"captcha"`
	Email     string `json:"email"`
}

func checkCaptcha(r *ReqCaptcha) (bool, error) {
	captchaId := r.CaptchaId
	captchaStr := r.Captcha
	ok := captcha.VerifyString(captchaId, captchaStr)
	if !ok {
		return false, errors.New("captcha verify failed")
	}
	return true, nil
}

func GetImageCaptcha(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()
	fileName, err := GetCaptchaFile()
	if err != nil {
		return
	}
	getBase64, err := GetBase64(fileName)
	if err != nil {
		return
	}
	split := strings.Split(fileName, ".")
	captchaId := split[0]
	respCaptcha := RespCaptcha{
		CaptchaId: captchaId,
		Base64:    getBase64,
	}
	c.JSON(200, gin.H{
		"code":    2000,
		"message": "get captcha success",
		"data":    respCaptcha,
	})
}
func EmailCaptcha(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()
	req := &ReqCaptcha{}
	if err = c.ShouldBindJSON(req); err != nil {
		return
	}
	//1. 校验用户图片验证码是否正确
	ok, err := checkCaptcha(req)
	if err != nil {
		return
	}
	if !ok {
		err = errors.New("captcha verify failed")
	}
	//2. 生成验证码，存入redis
	code := utils.RandNum(6)
	err = global.RedisClient.SetEMailCaptcha(req.Email, code)
	if err != nil {
		return
	}
	//3. 拿到验证码
	newEmail := email.NewEmail()
	newEmail.From = "1665400978@qq.com"
	newEmail.To = []string{req.Email}
	newEmail.Subject = "验证码"
	newEmail.Text = []byte("您的验证码是：" + code)
	// 向邮件channel中添加邮件
	global.EmailLists <- newEmail
}

// TODO swag
// ReloadCaptcha GetCaptcha 获取验证码
func ReloadCaptcha(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()
	s := struct {
		CaptchaId string `json:"captchaId"`
	}{}
	if err = c.ShouldBindJSON(&s); err != nil {
		return
	}
	err = reloadCaptcha(s.CaptchaId)
	if err != nil {
		return
	}
	getBase64, err := GetBase64(s.CaptchaId + ".png")
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":    2000,
		"message": "reload captcha success",
		"data": RespCaptcha{
			CaptchaId: s.CaptchaId,
			Base64:    getBase64,
		},
	})

}

func reloadCaptcha(captchaId string) error {
	file, err := os.Open(captchaId + ".png")
	if err != nil {
		return err
	}
	defer file.Close()
	ok := captcha.Reload(captchaId)
	if !ok {
		return errors.New("reload captcha failed")
	}
	return nil
}

func GetCaptchaFile() (string, error) {

	captchaId := captcha.NewLen(6)
	fileName := captchaId + ".png"
	f, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()
	err = captcha.WriteImage(f, captchaId, captcha.StdWidth, captcha.StdHeight)
	if err != nil {
		return "", err
	}
	return fileName, err

}

func GetBase64(fileName string) (string, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	bytes := make([]byte, base64.StdEncoding.EncodedLen(len(file)))
	base64.StdEncoding.Encode(bytes, file)
	s := string(bytes)
	return s, nil
}
