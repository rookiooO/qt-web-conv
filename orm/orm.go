package orm

import (
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Option struct {
	Driver          string `yaml:"driver"`
	DSN             string `yaml:"dsn"`
	MaxIdleConns    int    `yaml:"max-idle-conns"`
	MaxOpenConns    int    `yaml:"max-open-conns"`
	ConnMaxLifetime int    `yaml:"conn-max-lifetime"`
	ConnMaxIdleTime int    `yaml:"conn-max-idle-time"`
	LogMode         string `yaml:"log-mode"`
}

func New(opt *Option) (*gorm.DB, error) {
	var (
		db      *gorm.DB
		sqlDB   *sql.DB
		logMode logger.Interface
		err     error
	)

	if opt.LogMode == "info" {
		logMode = logger.Default.LogMode(logger.Info)
	} else {
		logMode = logger.Default.LogMode(logger.Error)
	}

	db, err = gorm.Open(postgres.Open(opt.DSN), &gorm.Config{
		Logger: logMode,
	})

	defer func(sqlDB *sql.DB) {
		if err != nil {
			_ = sqlDB.Close()
		}
	}(sqlDB)

	if err != nil {
		return nil, err
	}

	sqlDB, err = db.DB()
	if err != nil {
		return nil, err
	}

	if opt.MaxIdleConns != 0 {
		sqlDB.SetMaxIdleConns(opt.MaxIdleConns)
	}
	if opt.MaxOpenConns != 0 {
		sqlDB.SetMaxOpenConns(opt.MaxOpenConns)
	}
	if opt.ConnMaxLifetime != 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(opt.ConnMaxLifetime))
	}
	if opt.ConnMaxIdleTime != 0 {
		sqlDB.SetConnMaxIdleTime(time.Duration(opt.ConnMaxIdleTime))
	}

	return db, err
}
