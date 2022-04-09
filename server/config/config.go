package config

type Config struct {
	Port int `yaml:"port"`
	PG struct{
		Port int `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Db string `yaml:"db"`
	} `yaml:"pg"`
}