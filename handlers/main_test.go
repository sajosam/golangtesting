package handlers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var mainhadler Handler


func TestMain(m *testing.M) {
	dsn := "host=localhost user=postgres password=root dbname=forapi port=5433 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
	usrHandler.DB = db
    usrHandler.Connect("localhost", "postgres", "root", "forapi", "5433")

    code := m.Run()

    dbInstance, _ := usrHandler.DB.DB()
    dbInstance.Close()

    os.Exit(code)
}

func TestConnect(t *testing.T) {
	assert.NotNil(t, usrHandler.DB)
}

