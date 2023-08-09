// Package setting store all the configuration and constant of this project.
package setting

import (
	"flag"
	"github.com/spf13/viper"
	"os"
	"time"
)

type LogConfig struct {
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
		Jwt          struct {
			Secret           string        `mapstructure:"secret"`
			Expire           time.Duration `mapstructure:"expire"`
			RememberMeExp    time.Duration `mapstructure:"remember-me-expire"`
			RefreshBeforeExp time.Duration `mapstructure:"refresh-before-expire"`
		} `mapstructure:"jwt"`
		Cookie struct {
			Path   string `mapstructure:"path"`
			Domain string `mapstructure:"domain"`
			Secure bool   `mapstructure:"secure"`
		} `mapstructure:"cookie"`
		Csrf struct {
			ExcludePatterns []string `mapstructure:"exclude-patterns"`
		} `mapstructure:"csrf"`
	} `mapstructure:"server"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		DBName   string `mapstructure:"db_name"`
		User     string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		LogMode  string `mapstructure:"log_mode"`
	} `mapstructure:"database"`
	Redis struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
	} `mapstructure:"redis"`
	OSS struct {
		Endpoint        string `mapstructure:"endpoint"`
		AccessKeyId     string `mapstructure:"access-key-id"`
		AccessKeySecret string `mapstructure:"access-key-secret"`
		BucketName      string `mapstructure:"bucket-name"`
		SaveFolder      string `mapstructure:"save-folder"`
	} `mapstructure:"oss"`
	Mail struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		From     string `mapstructure:"from"`
		Template struct {
			Path           string `mapstructure:"path"`
			ReplySubject   string `mapstructure:"reply-subject"`
			ReviewSubject  string `mapstructure:"review-subject"`
			ReplyURLFormat string `mapstructure:"reply-url-format"`
			ReviewURL      string `mapstructure:"review-url"`
			AdminEmail     string `mapstructure:"admin-email"`
		} `mapstructure:"template"`
	} `mapstructure:"mail"`
	Log struct {
		Default LogConfig `mapstructure:"default"`
		Api     LogConfig `mapstructure:"api"`
		Error   LogConfig `mapstructure:"error"`
	} `mapstructure:"log"`
	Task struct {
		Cron struct {
			PersistArticleReadCount string `mapstructure:"persist-article-read-count"`
		} `mapstructure:"cron"`
	} `mapstructure:"task"`
	WordsFilter struct {
		WordsPath *string `mapstructure:"words-path"`
	} `mapstructure:"words-filter"`
}

var Config = config{}

const (
	ArticleReadCountPrefix = "bhs:article:read_count:"
	RecoveryAbortKey       = "bhs.recovery.abort"
	UnauthorizedMessage    = "请登录后再进行操作"
	InternalErrorMessage   = "内部错误，请联系管理员"
)

var (
	StatusArticlePublished = "PUBLISHED"
	StatusArticleDraft     = "DRAFT"
	StatusArticleHidden    = "HIDDEN"
	StatusCommentPublished = "PUBLISHED"
	StatusCommentReview    = "REVIEW"
	StatusCommentHidden    = "HIDDEN"
)

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
		panic(err.Error())
	}

	if err := v.Unmarshal(&Config); err != nil {
		panic(err.Error())
	}
}
