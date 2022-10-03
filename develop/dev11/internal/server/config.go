package server

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type Config struct {
	BindAddr         string `toml:"BIND_ADDR"`
	LogLevel         logrus.Level
	LogLevelString   string `toml:"LOG_LEVEL"`
	DatabaseHost     string `toml:"DATABASE_HOST"`
	DatabaseDBName   string `toml:"DATABASE_DB"`
	DatabaseUser     string `toml:"DATABASE_USER"`
	DatabasePassword string `toml:"DATABASE_PASSWORD"`
	DatabaseSSLMode  string `toml:"DATABASE_SSLMODE"`
}

func MakeConfigFromFile(path string) (Config, error) {
	config := Config{}

	if _, err := toml.DecodeFile(path, &config); err != nil {
		return config, err
	}

	var logLevels = map[string]logrus.Level{
		"PANIC": logrus.PanicLevel,
		"FATAL": logrus.FatalLevel,
		"ERROR": logrus.ErrorLevel,
		"WARN":  logrus.WarnLevel,
		"INFO":  logrus.InfoLevel,
		"DEBUG": logrus.DebugLevel,
		"TRACE": logrus.TraceLevel,
	}

	config.LogLevel = logLevels[config.LogLevelString]

	return config, nil
}
