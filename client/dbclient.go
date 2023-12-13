package client

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"etov/conf/db"
)

func ConnectDB(conf db.Mysql) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/etov?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port)
	client, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		logrus.Error(err.Error())
		return nil
	}
	return client
}
