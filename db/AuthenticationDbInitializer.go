package db

import (
	"go-auth/models/users"
	"go-auth/utils"
	"time"

	dbClient "github.com/12345debdut/go-sql/client"
	sqlConfigModel "github.com/12345debdut/go-sql/models"
	"gorm.io/gorm"
)

type AuthenticationDbInitializer struct{}

func (p *AuthenticationDbInitializer) InitDb() (*gorm.DB, error, func() error) {
	sqlDbClient := dbClient.NewSqlClient()
	dbModels := []interface{}{&users.User{}}
	goOrm, dbError := sqlDbClient.Connect(sqlConfigModel.SqlClientConnectionConfig{
		Host:       utils.GetEnv("DB_HOST", "localhost"),
		Port:       utils.GetEnv("DB_PORT", "5432"),
		User:       utils.GetEnv("DB_USER", "root"),
		Password:   utils.GetEnv("DB_PASSWORD", ""),
		DbName:     utils.GetEnv("DB_NAME", "go_auth"),
		SslMode:    sqlConfigModel.Disable,
		Models:     dbModels,
		RetryWait:  2 * time.Second,
		RetryCount: 10,
		OrmConfig:  &gorm.Config{},
	})
	return goOrm, dbError, func() error {
		return sqlDbClient.Disconnect()
	}
}
