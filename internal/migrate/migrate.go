package main

import (
	"github.com/H0lyDiv3r/go-headless-cms/internal/initializer"
	"github.com/H0lyDiv3r/go-headless-cms/internal/model"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectDB()
}

func main() {
	initializer.DB.AutoMigrate(&model.User{})
}
