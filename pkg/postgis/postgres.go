package postgresql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"surge/config"
)

var (
	db *gorm.DB = nil
)

func Init(cnf *config.Config) error {
	if db == nil {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			cnf.Database.Host,
			cnf.Database.User,
			cnf.Database.Password,
			cnf.Database.DbName,
			cnf.Database.Port)
		pdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}
		db = pdb
	}
	return nil
}

func Get() *gorm.DB {
	return db
}
