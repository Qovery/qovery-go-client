package main

import (
	"flag"
	"fmt"
	"github.com/Qovery/qovery-go-client"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var (
	configurationFilename = flag.String("config-filename", "../../test_files/local_configuration.json", "Qovery configuration filename")
	databaseName          = flag.String("dbname", "my-pql", "")
)

func main() {
	flag.Parse()
	qv, err := qovery.New(configurationFilename)
	if err != nil {
		log.Fatalf("fail to init qv client: %s", err.Error())
	}

	dbConf := qv.GetDatabaseConfigurationByName(*databaseName)
	if dbConf == nil {
		log.Fatalf("fail to get database name %s", *databaseName)
	}

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%d", dbConf.Host, dbConf.Username, dbConf.Name, dbConf.Password, dbConf.Port)
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("fail to connect to dbConf: %s", err.Error())
	}
	defer db.Close()
	log.Printf("connection to '%s' successful", dbConf.Name)
}
