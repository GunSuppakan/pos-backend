package database

import (
	"fmt"
	"pos-backend/internal/domain"
	"pos-backend/internal/infrastructure/logs"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var (
	DBConn *gorm.DB
)

type SqlLogger struct {
	logger.Interface
}

func InitDatabase() (*gorm.DB, error) {
	logs.Info("Connecting to database : " + viper.GetString("pg.host"))
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=require TimeZone=Asia/Bangkok",
		viper.GetString("pg.host"),
		viper.GetString("pg.username"),
		viper.GetString("pg.password"),
		viper.GetString("pg.name"),
		viper.GetString("pg.port"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("connect database failed: %w", err)
	}

	// ใช้ dbresolver หลังจาก db ไม่ nil แล้วเท่านั้น
	if err := db.Use(dbresolver.Register(dbresolver.Config{
		Sources:           []gorm.Dialector{postgres.Open(dsn)},
		Policy:            dbresolver.RandomPolicy{},
		TraceResolverMode: true,
	})); err != nil {
		return nil, fmt.Errorf("register dbresolver failed: %w", err)
	}

	if err := db.AutoMigrate(
		&domain.Account{},
		&domain.Product{},
		&domain.FilePath{},
		&domain.Category{},
		&domain.Stock{},
		&domain.StockTransaction{},
		&domain.Order{},
		&domain.OrderDetail{},
	); err != nil {
		return nil, fmt.Errorf("auto migrate failed: %w", err)
	}

	logs.Info("✅ Database connected")
	return db, nil
}
