package config

type RedisConfig struct {
	RedisHost     string `yaml:"host"`
	RedisPort     string `yaml:"port"`
	RedisPassword string `yaml:"password"`
}
