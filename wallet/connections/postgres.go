package connections

import (
	"fmt"
	"github.com/mhd7966/arvan/wallet/configs"
	"github.com/mhd7966/arvan/wallet/log"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	var err error
	c := configs.Cfg.Postgres
	DB, err = gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		c.Host,
		c.User,
		c.Pass,
		c.Name,
		c.Port,
	)), &gorm.Config{})

	if err != nil {
		log.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("Connect to DB Failed !!!!!")
		return err
	}
	
	return nil
}

func CloseDB() error {

	sqlDB, err := DB.DB()

	if err != nil {
		log.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("Get sqlDB Failed !!!!!")
		return err
	}

	sqlDB.Close()

	return nil
}
