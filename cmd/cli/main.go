package main

import (
	"flag"
	"log"

	"github.com/pkg/errors"

	"github.com/scriptted/goticker/internal/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	configFile = ""

	db   *gorm.DB
	conf *config.Config
)

func init() {
	flag.StringVar(&configFile, "config", configFile, "configuration file")
}

func main() {
	flag.Parse()

	var err error

	if configFile != "" {
		conf, err = config.NewFromFile(configFile)
		if err != nil {
			log.Fatalf("%+v", errors.Wrapf(err, "[Config] Could not load config file '%s'", configFile))
		}
	} else {
		panic("[Config] You must specify a config file path")
	}

	db, err = gorm.Open("sqlite3", "./data/goticker.db")
	if err != nil {
		panic("[DB] Connection failed")
	}

	defer db.Close()
}
