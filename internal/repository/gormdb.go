package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"template.com/restapi/internal/conf"
	"time"
)

type Repo struct {
	Db *gorm.DB
}

// NewConnection initialization of the connection
func NewConnection(connection string, conf *conf.AppConfig) (*Repo, error) {
	// here all the configuration for DB connection
	if connection == "" && conf == nil {
		return nil, fmt.Errorf("no connection provided")
	}

	var err error
	if connection == "" {
		connection = fmt.Sprintf("%s:%s@(%s:%s)/%s", conf.Db.User, conf.Db.Pass, conf.Db.URL, conf.Db.Port, conf.Db.Name)
	}
	dbGorm, err := gorm.Open(mysql.Open(connection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Silent, Error, Warn, Info
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := dbGorm.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &Repo{Db: dbGorm}, nil
}
