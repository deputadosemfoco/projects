package main

import (
	"fmt"
	"os"

	"github.com/deputadosemfoco/go-libs/redisdb"
	"github.com/deputadosemfoco/projects/repositories"
	"github.com/deputadosemfoco/projects/routes"
	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	godotenv.Load()

	in, _ := os.Open("banner.txt")
	defer in.Close()
	banner.Init(os.Stdout, true, false, in)

	e := routes.SetUp()
	port := os.Getenv("PORT")

	_redis, _ := redisdb.Connect()
	defer _redis.Close()

	repositories.Bootstrap()

	fmt.Println(fmt.Sprintf("projects reader service running on port %s", port))

	e.Run(standard.New(":" + port))
}
