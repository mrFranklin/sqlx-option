# sqlx-option
more convenient to init [golang-sqlx](https://github.com/jmoiron/sqlx) db

###Feature:
- init sqlx db by using `Option` type:
```
type Option struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     uint32 `mapstructure:"port"`
	Charset  string `mapstructure:"charset"`
	DbName   string `mapstructure:"dbname"`

	Driver   string `mapstructure:"driver"`  // sql driver for sqlx, the default is `mysql`
}
``` 
- init sqlx db by using yaml config file, format: the key is `mysql.dev`
```
mysql:
  dev:
    user: "root"
    password: ""
    host: "127.0.0.1"
    port: "3306"
    charset: "utf8mb4"
    dbname: eth-store
```

###examples:
init sqlx db by using `Option` type:
```
    opt := &Option{
        User: "root",
        Password: "",
        Host: "127.0.0.1",
        Port: 3306,
        Charset: "utf8mb4",
        DbName: "sqlx",
    }
    db, err := sqlxopt.NewDBWithOption(opt)
```

init sqlx db by using yaml config file, the key is `mysql.dev`
```
 
db, err := sqlxopt.NewDBWithConfigFile("config.yaml", "mysql.dev")

```
config.yaml:
```
mysql:
  dev:
    User: "root"
    Password: ""
    Host: "127.0.0.1"
    Port: 3306
    Charset: "utf8mb4"
    DbName: "sqlx"
```
