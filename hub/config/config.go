package config

import (
	"errors"
	"flag"
	"fmt"
	"github.com/suixinio/headscale-hub/util"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

// 系统配置，对应yml
// viper内置了mapstructure, yml文件用"-"区分单词, 转为驼峰方便

// 全局配置变量
var Conf = new(config)

type config struct {
	System    *SystemConfig    `mapstructure:"system" json:"system"`
	Logs      *LogsConfig      `mapstructure:"logs" json:"logs"`
	Database  *DatabaseConfig  `mapstructure:"database" json:"database"`
	Casbin    *CasbinConfig    `mapstructure:"casbin" json:"casbin"`
	Jwt       *JwtConfig       `mapstructure:"jwt" json:"jwt"`
	RateLimit *RateLimitConfig `mapstructure:"rate-limit" json:"rateLimit"`
	Headscale *HeadscaleConfig `mapstructure:"headscale" json:"headscale"`
}

// 设置读取配置信息
func InitConfig() {
	confPath := ""
	flag.StringVar(&confPath, "c", "./config.yml", "Path to the config file.")
	flag.Parse()
	viper.SetDefault("config", confPath)
	configFile := viper.GetString("config")
	if len(configFile) <= 0 {
		panic(fmt.Errorf("config path is empty"))
	}

	if fileInfo, err := os.Stat(configFile); err != nil || fileInfo.IsDir() {
		if errors.Is(err, os.ErrNotExist) {
			panic(fmt.Errorf("config path not exist"))
		}
		panic(fmt.Errorf("config file is error or not dir"))
	}
	viper.SetConfigFile(configFile)
	// 读取配置信息
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Failed to read configuration file:%s \n", err))
	}

	// 热更新配置
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 将读取的配置信息保存至全局变量Conf
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("初始化配置文件失败:%s \n", err))
		}
		// 读取rsa key
		Conf.System.RSAPublicBytes = util.RSAReadKeyFromFile(Conf.System.RSAPublicKey)
		Conf.System.RSAPrivateBytes = util.RSAReadKeyFromFile(Conf.System.RSAPrivateKey)
	})

	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("初始化配置文件失败:%s \n", err))
	}
	// 读取rsa key
	Conf.System.RSAPublicBytes = util.RSAReadKeyFromFile(Conf.System.RSAPublicKey)
	Conf.System.RSAPrivateBytes = util.RSAReadKeyFromFile(Conf.System.RSAPrivateKey)

}

type SystemConfig struct {
	Mode            string `mapstructure:"mode" json:"mode"`
	UrlPathPrefix   string `mapstructure:"url-path-prefix" json:"urlPathPrefix"`
	Host            string `mapstructure:"host" json:"host"`
	Port            int    `mapstructure:"port" json:"port"`
	InitData        bool   `mapstructure:"init-data" json:"initData"`
	RSAPublicKey    string `mapstructure:"rsa-public-key" json:"rsaPublicKey"`
	RSAPrivateKey   string `mapstructure:"rsa-private-key" json:"rsaPrivateKey"`
	RSAPublicBytes  []byte `mapstructure:"-" json:"-"`
	RSAPrivateBytes []byte `mapstructure:"-" json:"-"`
}

type LogsConfig struct {
	Level      zapcore.Level `mapstructure:"level" json:"level"`
	Path       string        `mapstructure:"path" json:"path"`
	MaxSize    int           `mapstructure:"max-size" json:"maxSize"`
	MaxBackups int           `mapstructure:"max-backups" json:"maxBackups"`
	MaxAge     int           `mapstructure:"max-age" json:"maxAge"`
	Compress   bool          `mapstructure:"compress" json:"compress"`
}

type DatabaseConfig struct {
	Type     string         `mapstructure:"type" json:"type"`
	Debug    bool           `mapstructure:"debug" json:"debug"`
	Gorm     GormConfig     `mapstructure:"gorm" json:"gorm"`
	Sqlite   SqliteConfig   `mapstructure:"sqlite" json:"sqlite"`
	Postgres PostgresConfig `mapstructure:"postgres" json:"postgres"`
}

type GormConfig struct {
	SkipErrRecordNotFound bool   `mapstructure:"skip_err_record_not_found" json:"skip_err_record_not_found"`
	TablePrefix           string `mapstructure:"table_prefix" json:"table_prefix"`
}

type SqliteConfig struct {
	Path          string `mapstructure:"path" json:"path"`
	WriteAheadLog bool   `mapstructure:"write_ahead_log" json:"write_ahead_log"`
}

type PostgresConfig struct {
	Host                string `mapstructure:"host" json:"host"`
	Port                int    `mapstructure:"port" json:"port"`
	Name                string `mapstructure:"name" json:"name"`
	User                string `mapstructure:"user" json:"user"`
	Pass                string `mapstructure:"pass" json:"pass"`
	MaxOpenConnections  int    `mapstructure:"max_open_conns" json:"max_open_conns"`
	MaxIdleConnections  int    `mapstructure:"max_idle_conns" json:"max_idle_conns"`
	ConnMaxIdleTimeSecs int    `mapstructure:"conn_max_idle_time_secs" json:"conn_max_idle_time_secs"`
	Ssl                 string `mapstructure:"ssl" json:"ssl"`
}

type CasbinConfig struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath"`
}

type JwtConfig struct {
	Realm      string `mapstructure:"realm" json:"realm"`
	Key        string `mapstructure:"key" json:"key"`
	Timeout    int    `mapstructure:"timeout" json:"timeout"`
	MaxRefresh int    `mapstructure:"max-refresh" json:"maxRefresh"`
}

type RateLimitConfig struct {
	FillInterval int64 `mapstructure:"fill-interval" json:"fillInterval"`
	Capacity     int64 `mapstructure:"capacity" json:"capacity"`
}

type HeadscaleConfig struct {
	Host   string `mapstructure:"host" json:"host"`
	ApiKey string `mapstructure:"api_key" json:"api_key"`
}
