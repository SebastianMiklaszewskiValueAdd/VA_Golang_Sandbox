package main

import (
	"ToDoApp/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	infrastructure.Setup(r)

	return r
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Environment configuration file not found!")
	}

	r := setupRouter()
	err = r.Run(os.Getenv("SERVER_PORT"))
	if err != nil {
		panic("Server crashed!!! ")
	}
}
