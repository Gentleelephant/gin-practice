package config

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Config struct {
	Mysql       *MysqlConf   `yaml:"mysql"`
	Server      *Server      `yaml:"server"`
	LDAP        *LdapConf    `yaml:"ldap"`
	Redis       *RedisConfig `yaml:"redis"`
	EmailConfig *EmailConfig `yaml:"email"`
}
