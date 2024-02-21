package manager

import (
	"database/sql"
	"fmt"

	"enigmacamp.com/enigma-laundry-apps/config"
	_ "github.com/go-sql-driver/mysql"
)

type InfraManager interface {
	Conn() *sql.DB
	GetConfig() *config.Config
}

type infraManager struct {
	db  *sql.DB
	cfg *config.Config
}

func (i *infraManager) initDb() error {
	var dbConf = i.cfg.DbConfig
	// dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	dbConf.Host,
	// 	dbConf.Port,
	// 	dbConf.User,
	// 	dbConf.Password,
	// 	dbConf.Name)
	dataSourceName := fmt.Sprintf("%s:%s@/%s",
		dbConf.User,
		dbConf.Password,
		dbConf.Name)
	db, err := sql.Open(dbConf.Driver, dataSourceName)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	i.db = db
	return nil
}

func (i *infraManager) Conn() *sql.DB {
	return i.db
}

func (i *infraManager) GetConfig() *config.Config {
	return i.cfg
}

func NewInfraManager(configParam *config.Config) (InfraManager, error) {
	infra := &infraManager{
		cfg: configParam,
	}

	err := infra.initDb()
	if err != nil {
		return nil, err
	}
	return infra, nil
}

// DEPENDENCY INJECTION
