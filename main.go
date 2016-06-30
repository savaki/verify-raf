package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
)

type Options struct {
	Port string
}

var opts Options

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{"PORT", "3000", "port to run on", "PORT", &opts.Port, false},
	}
	app.Action = Run
	app.Run(os.Args)
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": strconv.FormatInt(time.Now().UnixNano(), 10),
	})
}

func Verify(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"updateSuccessful":           true,
		"externalMarketingPartnerId": 1234,
	})
}

func Run(c *cli.Context) error {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.POST("/api/login", Login)
	router.POST("/api/verify", Verify)
	err := router.Run(":" + opts.Port)
	fmt.Println(err)
	return err
}
