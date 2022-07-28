package config

import (
	"github.com/spf13/viper"
	"time"
)

type (
	Config struct {
		Environment string
		HTTP        HTTPConfig     `mapstructure:"http"`
		Postgres    PostgresConfig `mapstructure:"postgres"`
		Jwt         JwtConfig
	}

	HTTPConfig struct {
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderMegaBytes"`
	}
	PostgresConfig struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		SSLMode  string `mapstructure:"sslmode"`
	}
	JwtConfig struct {
		Salt string
	}
)

// Initialize config using config in the configsDir
// and environment variables
func Init(configsDir, environment, envFilePath string) (*Config, error) {
	fillDefaults()

	if err := parseEnv(envFilePath); err != nil {
		return nil, err
	}

	if err := parseConfigs(configsDir, environment); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	cfg.Environment = environment

	setFromEnv(&cfg)

	return &cfg, nil
}

func parseConfigs(configsDir, configName string) error {
	viper.AddConfigPath(configsDir)
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}
