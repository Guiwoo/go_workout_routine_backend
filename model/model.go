package model

import (
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_USER     = "postgres"
	DB_PORT     = 5432
	DB_PASSWORD = "123"
	DB_NAME     = "guiwoopark"
)

var (
	en  *xorm.Engine
	err error
)

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	en, err = xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		log.Panic("Engine creation fialed", err)
	}
}

func db_sync() {
	en.Sync(new(User_Type))
}

func DB_Connect() {
	db_sync()
}

func Close() {
	en.Close()
}

func DB_Handler() *xorm.Engine {
	return en
}
