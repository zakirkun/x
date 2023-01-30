package database

import (
	"errors"
	"fmt"

	"github.com/zakirkun/x/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(opt DBModel) IDatabase {
	return DBModel{
		Driver:   opt.Driver,
		Host:     opt.Host,
		Port:     opt.Port,
		Name:     opt.Name,
		Username: opt.Username,
		Password: opt.Password,
	}
}

func (i DBModel) InitDB() (*gorm.DB, *error) {
	var connectionUrl string

	switch i.Driver {
	case "mysql":
		connectionUrl = fmt.Sprintf(MYSQL_CONFIG, i.Username, i.Password, i.Host, i.Port, i.Name)
		db, err := gorm.Open(mysql.Open(connectionUrl), &gorm.Config{})
		if err != nil {
			logger.Warn(fmt.Sprintf("Failed connect to database : %v", err))
			return nil, &err
		}

		return db, nil
	case "postgresql":
		connectionUrl = fmt.Sprintf(POSGRES_CONFIG, i.Username, i.Password, i.Name, i.Host, i.Port, "disable")
		db, err := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{})
		if err != nil {
			logger.Warn(fmt.Sprintf("Failed connect to database : %v", err))
			return nil, &err
		}

		return db, nil
	default:
		logger.Warn("Please select driver.")
		err := errors.New("driver not selected/invalid")
		return nil, &err
	}

}
