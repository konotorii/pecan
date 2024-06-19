package pecan

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var port, port_exists = strconv.Atoi(os.Getenv("PORT"))

var git_token = os.Getenv("GITHUB_TOKEN")

var git_repo = os.Getenv("GITHUB_REPO")

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(color.Ize(color.White, "No .env was found!"))
	}
}

func main() {
	fmt.Print(port)
	r := engine()
	r.Use(gin.Logger())
	if err := engine().Run(":8080"); err != nil {
		log.Fatal("unable to start:", err)
	}
}

func engine() *gin.Engine {
	r := gin.New()

	return r
}
