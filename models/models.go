package models

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/lunny/xorm"
)

var DBConfig struct {
	SslMode, Name, User string
}

func LoadDBConfig() {
	DBConfig.SslMode = os.Getenv("DB_SSLMODE")
	DBConfig.Name = os.Getenv("DB_NAME")
	DBConfig.User = os.Getenv("DB_USER")
}

func InitOrm() *xorm.Engine {
	LoadDBConfig()
	orm, err := xorm.NewEngine("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=%s",
		DBConfig.User, DBConfig.Name, DBConfig.SslMode))
	if err != nil {
		panic(err)
	}
	return orm
}
