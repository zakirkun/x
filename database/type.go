package database

import "gorm.io/gorm"

type DBModel struct {
	Driver   string
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}

type IDatabase interface {
	InitDB() (*gorm.DB, *error)
}

const (
	POSGRES_CONFIG = "user=%s password=%s dbname=%s host=%s port=%s sslmode=%s"
	MYSQL_CONFIG   = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
)
