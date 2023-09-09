package main

import (
	"go-app/inits"
	"go-app/models"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	inits.DB.AutoMigrate(&models.User{})
	inits.DB.AutoMigrate(&models.Post{})
}
