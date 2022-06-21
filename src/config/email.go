package config

type EmailConfig struct {
	Address  string `yaml:"address"`  // eg: smtp.qq.com:25
	Count    int    `yaml:"count"`    // 邮件池大小
	Username string `yaml:"username"` // 发件人邮箱
	Password string `yaml:"password"` // 发件人邮箱密码
	Host     string `yaml:"host"`     // 发件人邮箱服务器 eg: smtp.qq.com
}
