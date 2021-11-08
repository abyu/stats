package internal

import (
	"gorm.io/gorm"
	"time"
)

//TokenBalance ...
type TokenBalance struct {
	Balance float64
	TokenName string
	TimeStamp time.Time
}

//MetricsStore ...
type MetricsStore interface {
	New(balance TokenBalance)
}

type metricsStore struct {
	DB *gorm.DB
}

//Migrate ...
func (m *metricsStore) Migrate() error {
	return m.DB.AutoMigrate(&TokenBalance{})
}

//NewMetricsStoreMigrator ...
func NewMetricsStoreMigrator(DB *gorm.DB) Migrator {
	return &metricsStore{DB: DB}
}

//NewMetricsStore ...
func NewMetricsStore(DB *gorm.DB) MetricsStore {
	return &metricsStore{DB: DB}
}

//New creates a new entity
func (m *metricsStore) New(balance TokenBalance) {
	m.DB.Create(balance)
}
