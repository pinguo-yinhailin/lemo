package model

type SettingConfig struct {
	Name    string   `mapstructure:"name"`
	IsAdmin bool     `mapstructure:"isAdmin"`
	MaxPage int      `mapstructure:"maxPage"`
	Pi      float64  `mapstructure:"pi"`
	Emails  []string `mapstructure:"emails"`
}
