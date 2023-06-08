package depedencies

import (
	"fmt"
	"go-web-boilerplate/shared/config"
	"go-web-boilerplate/shared/dto"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(env *config.EnvConfig, log *logrus.Logger) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		env.DBHost,
		env.DBUser,
		env.DBPassword,
		env.DBName,
		env.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Errorf("failed to connect to database, with error: %s", err.Error())
	}

	setConnectionConfiguration(db)

	log.Infoln("connected to databse with configuration: %s", dsn)

	migrateSchema(db, log)

	return db
}

func setConnectionConfiguration(db *gorm.DB) {
	postgresDb, _ := db.DB()
	postgresDb.SetMaxIdleConns(10)
	postgresDb.SetMaxOpenConns(100)
	postgresDb.SetConnMaxLifetime(time.Hour)
}

func migrateSchema(db *gorm.DB, log *logrus.Logger) {
	err := db.AutoMigrate(
		&dto.User{},
		&dto.PasswordReset{},
	)

	if err != nil {
		log.Errorf("error migrating schema, err: %s", err.Error())
	}

	log.Infoln("database migrated")
}