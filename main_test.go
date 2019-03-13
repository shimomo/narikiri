package main

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	code := m.Run()
	os.Exit(code)
}
