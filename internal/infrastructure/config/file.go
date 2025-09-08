package config

import "github.com/caarlos0/env/v6"

type Config struct {
	App struct {
		Key string `env:"APP_KEY"`
		Env string `env:"APP_ENV"`
	}
	WebServer struct {
		REST struct {
			Host string `env:"WEBSERVER_REST_HOST"`
			Port string `env:"WEBSERVER_REST_PORT"`
		}
	}
	Logger struct {
	}
	Database struct {
		Driver   string `env:"DB_DRIVER"`
		Name     string `env:"DB_NAME"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT"`
	}
}

func NewConfig() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c *Config) GetWebServerRestHost() string {
	return c.WebServer.REST.Host
}

func (c *Config) GetWebServerRestPort() string {
	return c.WebServer.REST.Port
}

func (c *Config) GetAppKey() string {
	return c.App.Key
}

func (c *Config) GetDbDriver() string {
	return c.Database.Driver
}

func (c *Config) GetDbHost() string {
	return c.Database.Host
}

func (c *Config) GetDbPort() string {
	return c.Database.Port
}

func (c *Config) GetDbName() string {
	return c.Database.Name
}

func (c *Config) GetDbUser() string {
	return c.Database.User
}

func (c *Config) GetDbPassword() string {
	return c.Database.Password
}
