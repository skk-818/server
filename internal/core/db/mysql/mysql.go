package mysql

import (
	"fmt"
	"server/internal/core/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SystemDB 系统数据库类型
type SystemDB struct {
	*gorm.DB
}

// ImDB IM数据库类型
type ImDB struct {
	*gorm.DB
}

func newMySQL(cfg *config.Mysql) (*gorm.DB, error) {
	if cfg.Dbname == "" {
		return nil, fmt.Errorf("数据库名不能为空")
	}

	dsn := cfg.DSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("数据库连接失败: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

	return db, nil
}

func NewSystemDB(cfg *config.Mysql) (*SystemDB, error) {
	db, err := newMySQL(cfg)
	if err != nil {
		return nil, err
	}
	return &SystemDB{DB: db}, nil
}

func NewImDB(cfg *config.Mysql) (*ImDB, error) {
	db, err := newMySQL(cfg)
	if err != nil {
		return nil, err
	}
	return &ImDB{DB: db}, nil
}
