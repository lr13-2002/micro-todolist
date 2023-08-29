package dao

import (
	"context"
	"fmt"
	"micro-todolist/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var _db *gorm.DB

func InitDB() {
	host := config.DbHost
	port := config.DbPort
	user := config.DbUser
	password := config.DbPassWord
	database := config.DbName
	charset := config.Charset
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", user, password, host, port, database, charset)
	fmt.Println(dsn)

	err := Database(dsn)
	if err != nil {
		fmt.Println(err)
	}
}

func Database(connString string) error {
	var ormLogger logger.Interface = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connString,
		DefaultStringSize:         256,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
		DisableDatetimePrecision:  true,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	_db = db
	migration()
	return err
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
