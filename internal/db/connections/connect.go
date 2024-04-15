package connections

import (
	"fmt"
	"github.com/YoungGoofy/quiz/internal/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	cfg, err := configs.ReadConf()
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.DatabaseName, cfg.DB.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
