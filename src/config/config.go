package config

type Config struct {
	ApiToken string `yaml:"api_token"`
	RedisUrl string `yaml:"redis_url"`
}
