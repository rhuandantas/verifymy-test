package repo

import (
	"database/sql"
	"fmt"
	"github.com/rhuandantas/verifymy-test/internal/config"
	"github.com/rhuandantas/verifymy-test/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConnection interface {
	GetDB() *gorm.DB
}

type MysqlORMConnection struct {
	db     *gorm.DB
	config config.ConfigProvider
}

func (conn MysqlORMConnection) GetDB() *gorm.DB {
	return conn.db
}

func createConnection(config config.ConfigProvider) (*sql.DB, error) {
	var user, password string
	host := config.GetString("db.mysql.host")
	port := config.GetString("db.mysql.port")
	database := config.GetString("db.mysql.database")
	user = config.GetEnv(config.GetString("db.mysql.user-key"))
	password = config.GetEnv(config.GetString("db.mysql.password-key"))
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewMysqlORMConn(config config.ConfigProvider) (DBConnection, error) {
	db, err := createConnection(config)
	if err != nil {
		return nil, err
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	err = gormDB.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	return &MysqlORMConnection{
		db: gormDB,
	}, nil
}
