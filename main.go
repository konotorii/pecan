package main

import (
	"fmt"
	"log"
	"pecan/util"

	"github.com/TwiN/go-color"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(color.Ize(color.White, "No .env was found!"))
	}

	util.ConfigInit()
}

func main() {
	r := engine()
	r.Use(gin.Logger())
	if err := engine().Run(fmt.Sprintf(":%d", util.Config.Port)); err != nil {
		log.Fatal("unable to start:", err)
	}
}

func engine() *gin.Engine {
	r := gin.New()

	r.GET("/download", download)
	r.GET("/download/:platform", downloadPlatform)
	r.GET("/update/:platform/:version", updatePlatformVersion)

	return r
}

func download(c *gin.Context) {

}

func downloadPlatform(c *gin.Context) {
	platform := c.Param("platform")

	if platform == "mac" {
		platform = "dmg"
	}
	if platform == "mac_arm64" {
		platform = "dmg_arm64"
	}

}

func updatePlatformVersion(c *gin.Context) {

}
