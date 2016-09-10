package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/engine/standard"
	"github.com/vitorsalgado/la-democracia/lib/go/redisdb"
	"github.com/vitorsalgado/la-democracia/projects-reader/routes"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("%s : %s", "Error loading .env file", err)
	}

	router := routes.Router{}
	e := router.SetUp()
	port := os.Getenv("Projects_Reader_Port")

	_redis, _ := redisdb.Connect()
	defer _redis.Close()

	fmt.Println(fmt.Sprintf("projects reader service running on port %s", port))

	e.Run(standard.New(":" + port))
}
