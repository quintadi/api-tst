package data

import (
	"fmt"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type User struct {
	Id       int64
	Name     string
	Email    string
	Password string
}

func CreateDBEngien() (*xorm.Engine, error) {
	connectionInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "127.0.0.1", 5432, "postgres", "p123", "authServer")
	engine, err := xorm.NewEngine("postgres", connectionInfo)

	fmt.Println(engine)
	if err == nil {
		return nil, err
	}
	if err := engine.Ping(); err != nil {
		return nil, err
	}
	if err := engine.Sync(new(User)); err != nil {
		return nil, err
	}
	return engine, err
}
