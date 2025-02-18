package config

type Config struct {
	System   System   `yaml:"system"`
	Mysql    Mysql    `yaml:"mysql"`
	Jwt      Jwt      `yaml:"jwt"`
	AliOSS   AliOSS   `yaml:"ali_oss"`
	BaiDuMap BaiDuMap `yaml:"baidu_map"`
}
