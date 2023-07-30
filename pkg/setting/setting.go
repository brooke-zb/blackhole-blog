package setting

import (
	"flag"
	"github.com/spf13/viper"
	"os"
)

type logConfig struct {
	Writer  string  `mapstructure:"writer"`
	Encoder string  `mapstructure:"encoder"`
	File    *string `mapstructure:"file"`
	Level   string  `mapstructure:"level"`
}

type config struct {
	Server struct {
		Host         string   `mapstructure:"host"`
		Port         int      `mapstructure:"port"`
		Proxy        []string `mapstructure:"proxy"`
		ProxyHeaders []string `mapstructure:"proxy-headers"`
	} `mapstructure:"server"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		DBName   string `mapstructure:"db_name"`
		User     string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		LogMode  string `mapstructure:"log_mode"`
	} `mapstructure:"database"`
	Log struct {
		Default logConfig `mapstructure:"default"`
		Api     logConfig `mapstructure:"api"`
		Error   logConfig `mapstructure:"error"`
	} `mapstructure:"log"`
	Task struct {
		Cron struct {
			PersistArticleReadCount string `mapstructure:"persist-article-read-count"`
		} `mapstructure:"cron"`
	} `mapstructure:"task"`
}

var Config = config{}

func readConfigPath() (path string) {
	// flag
	flag.StringVar(&path, "config-path", "", "choose config file path.")
	flag.Parse()
	if path != "" {
		return path
	}

	// env
	path, ok := os.LookupEnv("BH_BLOG_CONFIG_PATH")
	if ok {
		return path
	}

	// default
	return "./conf/config.toml"
}

func Setup() {
	v := viper.New()
	v.SetConfigFile(readConfigPath())
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&Config); err != nil {
		panic(err)
	}
}
