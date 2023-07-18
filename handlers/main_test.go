package handlers

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var mainhadler Handler


func TestMain(m *testing.M) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST_TEST"), os.Getenv("DB_USER_TEST"), os.Getenv("DB_PASSWORD_TEST"), os.Getenv("DB_NAME_TEST"), os.Getenv("DB_PORT_TEST"))
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
	mainhadler.DB = db
    mainhadler.Connect(os.Getenv("DB_HOST_TEST"), os.Getenv("DB_USER_TEST"), os.Getenv("DB_PASSWORD_TEST"), os.Getenv("DB_NAME_TEST"), os.Getenv("DB_PORT_TEST"))

    code := m.Run()

    dbInstance, _ := mainhadler.DB.DB()
    dbInstance.Close()

    os.Exit(code)
}

func TestConnect(t *testing.T) {
	assert.NotNil(t, mainhadler.DB)
}


