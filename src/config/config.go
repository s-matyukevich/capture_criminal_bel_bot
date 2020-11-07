package config

type SourceChat struct {
	Id   int64  `yaml:"id"`
	Name string `yaml:"name"`
	City string `yaml:"city"`
}

type Config struct {
	ApiToken           string        `yaml:"api_token"`
	RedisUrl           string        `yaml:"redis_url"`
	SourceChatIds      []*SourceChat `yaml:"source_chats"`
	SourceByTagChatIds []*SourceChat `yaml:"source_by_tag_chats"`
	AdminChats         []int         `yaml:"admin_chats"`
	TelegramApiId      int32         `yaml:"telegram_api_id"`
	TelegramApiHash    string        `yaml:"telegram_api_hash"`
	GeocodingKey       string        `yaml:"geocoding_key"`
}
