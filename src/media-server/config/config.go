package config

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
		Port    string `yaml:"port"`
		Timeout int    `yaml:"timeout"`
	}

	Storage struct {
		Local
	}

	Local struct {
		Path string `yaml:"path"`
	}
)
