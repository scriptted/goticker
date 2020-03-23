package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
	"github.com/scriptted/goticker/internal/config"
	"github.com/scriptted/goticker/internal/route"
	"gitlab.com/wpetit/goweb/middleware/container"
)

//nolint: gochecknoglobals
var (
	configFile = ""
	dbFile     = ""
	workdir    = ""
)

//nolint: gochecknoinits
func init() {
	flag.StringVar(&configFile, "config", configFile, "configuration file")
	flag.StringVar(&dbFile, "database", dbFile, "database file")
	flag.StringVar(&workdir, "workdir", workdir, "working directory")
}

func main() {
	flag.Parse()

	// Switch to new working directory if defined
	if workdir != "" {
		if err := os.Chdir(workdir); err != nil {
			log.Fatalf("%+v", errors.Wrapf(err, "[OS] Could not change working directory to '%s'", workdir))
		}
	}

	// Load configuration file if defined, dump and use default configuration otherwise
	var conf *config.Config

	var err error

	if configFile != "" {
		conf, err = config.NewFromFile(configFile)
		if err != nil {
			log.Fatalf("%+v", errors.Wrapf(err, "[Config] Could not load config file '%s'", configFile))
		}
	} else {
		conf = config.CreateDefault()
	}

	if dbFile == "" {
		log.Fatalf("%+v", errors.Wrapf(err, "[Database] Please provide a database file"))
	}

	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatalf("%+v", errors.Wrap(err, "[DB] Could not set database"))
	}
	defer db.Close()

	// Create service container
	ctn, err := getServiceContainer(conf, db)
	if err != nil {
		log.Fatalf("%+v", errors.Wrap(err, "[GOWEB] Could not create service container"))
	}

	r := chi.NewRouter()

	// Define base middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Expose service container on router
	r.Use(container.ServiceContainer(ctn))

	// Define routes
	if err := route.Mount(r, conf); err != nil {
		log.Fatalf("%+v", errors.Wrap(err, "[HTTP] Could not mount http routes"))
	}

	log.Printf("listening on '%s'", conf.HTTP.Address)
	if err := http.ListenAndServe(conf.HTTP.Address, r); err != nil {
		log.Fatalf("%+v", errors.Wrapf(err, "[HTTP] Could not listen on '%s'", conf.HTTP.Address))
	}
}
