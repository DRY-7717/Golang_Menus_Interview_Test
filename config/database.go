package config

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (config *Config) ConnectionPostgres() (*Postgres, error) {

	dbConnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.Psql.User,
		config.Psql.Password,
		config.Psql.Host,
		config.Psql.Port,
		config.Psql.Name,
	)

	db, err := gorm.Open(postgres.Open(dbConnection), &gorm.Config{})

	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgres] Failed to connect to database -1 " + config.Psql.Host)
		return nil, err
	}

	sqlDb, err := db.DB()

	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgres] Failed to connect to database - 2 ")
		return nil, err
	}

	sqlDb.SetMaxOpenConns(config.Psql.MaxOpen)
	sqlDb.SetMaxIdleConns(config.Psql.MaxIdle)

	return &Postgres{
		DB: db,
	}, nil

}
