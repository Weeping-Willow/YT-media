package config

import "time"

type (
	Config struct {
		App
		Http
		Storage
	}

	App struct {
		Name string `yaml:"name"`
	}

	Http struct {
		Port             string        `yaml:"port"`
		KeepAliveTimeout time.Duration `yaml:"keep_alive_timeout"`
	}

	Storage struct {
		Local
	}

	Local struct {
		Path string `yaml:"path"`
	}
)
