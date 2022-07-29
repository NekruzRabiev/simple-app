package config

import (
	"github.com/spf13/viper"
	"time"
)

var (
	defaults = map[string]interface{}{
		"http.port":                 "8000",
		"http.max_header_megabytes": 1,
		"http.readTimeout":          5 * time.Second,
		"http.writeTimeout":         5 * time.Second,
	}
)

func fillDefaults() {
	for k, v := range defaults {
		viper.SetDefault(k, v)
	}
}
