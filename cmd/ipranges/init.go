package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/nihab70/cloudorama/cloudkit/aws"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/revel/config"
)

var configFile = "app.conf"

func initDB() {
	log.Infof("Read config from file: %v", configFile)

	c, err := config.ReadDefault(configFile)
	if err != nil {
		log.Warnf("failed to read config file %v: %v", configFile, err.Error())
	}

	dbdriver, err := c.String("dev", "db.driver")
	if err != nil {
		log.Warnf("failed to read config for db driver: %v", err.Error())
	}

	dbfile, err := c.String("dev", "db.database")
	if err != nil {
		log.Warnf("failed to read config for db params: %v", err.Error())
	}
	log.Infof("Use DB connection %v, %v", dbdriver, dbfile)

	db, err := gorm.Open(dbdriver, dbfile)
	if err != nil {
		log.Fatalf("failed to connect database %v, %v: %v", dbdriver, dbfile, err.Error())
	}

	defer db.Close()
	aws.PrepareDB(db)

}
