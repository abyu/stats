package internal

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

//AppContext ...
type AppContext struct {
	DB *gorm.DB
	DBMigration Migrations
	FetchTask FetchTask
}

//BuildContext builds it
func BuildContext() (*AppContext, error){
	config, err := AppConfig()
	if err != nil {
		return nil, err
	}
	appCtx := &AppContext{}
	db, err := newDbConnection(*config)
	if err != nil {
		return nil, err
	}
	appCtx.DB = db
	appCtx.DBMigration = NewDBMigration(NewMetricsStoreMigrator(db))
	appCtx.FetchTask = NewFetchTask(NewMetricsStore(db), 30 * time.Minute)
	return appCtx, nil
}

func newDbConnection(config Config) (*gorm.DB, error) {
	log.Infof("Attempting to connect to %s:%s with user %s", config.DBHost, config.DBPort, config.DBUserName)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Australia/Sydney", config.DBHost, config.DBUserName, config.DBPassword, config.DB, config.DBPort)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
