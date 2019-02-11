package sqlxopt

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
)

type Option struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     uint32 `mapstructure:"port"`
	Charset  string `mapstructure:"charset"`
	DbName   string `mapstructure:"dbname"`

	Driver   string `mapstructure:"driver"`  // sql driver for sqlx, the default is `mysql`
}

func (opt Option) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local", opt.User, opt.Password, opt.Host, opt.Port, opt.DbName, opt.Charset)
}

// NewDB creates sqlx.DB by using `source` string and driver
func NewDB(driver, source string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(driver, source)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// NewDBWithOption creates sqlx.DB by using `Option`
func NewDBWithOption(opt Option) (*sqlx.DB, error) {
	if opt.Driver == "" {
		opt.Driver = "mysql"
	}
	return NewDB(opt.Driver, opt.String())
}

// NewDBWithConfig creates sqlx.DB by using yaml config file
func NewDBWithConfigFile(configFile string, key string) (*sqlx.DB, error) {
	opt, err := LoadOptionFromConfigFile(configFile, key)
	if err != nil {
		return nil, err
	}
	return NewDBWithOption(opt)
}

/*
LoadOptionFromConfigFile loads the mysql config from yaml file,
yaml config file example: (the key is mysql.dev)
mysql:
  dev:
    user: "root"
    password: ""
    host: "127.0.0.1"
    port: 3306
    charset: "utf8mb4"
    dbname: "sqlx"
 */

func LoadOptionFromConfigFile(configFile string, key string) (opt Option, err error) {
	vp := viper.New()
	vp.SetConfigFile(configFile)
	if err = vp.ReadInConfig(); err != nil {
		log.Fatalf("Can't read config file: %v", err)
		return opt, err
	}

	if err = vp.UnmarshalKey(key, &opt); err != nil {
		log.Fatalf("Can't unmarshal key values: %v", err)
	}
	return opt, err
}