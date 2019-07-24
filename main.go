package main

import (
	// "crypto/tls"
	// "crypto/x509"
	"github.com/gin-gonic/gin"
	"log"
	// "io/ioutil"
	// "net"
	"net/http"
	"os"
	// model "github.com/mikequentel/go-rest-demo/model"
	// "strconv"
	// "time"
)

var (
	SERVER_CERT string
	SERVER_KEY  string
	SERVER_PORT string
)

func init() {
	SERVER_CERT = os.Getenv("SERVER_CERT")
	SERVER_KEY = os.Getenv("SERVER_KEY")
	SERVER_PORT = os.Getenv("SERVER_PORT")
}

func root(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code":    http.StatusOK,
			"message": "go-rest-demo server",
		},
	)
}

func main() {
	r := gin.Default()
	r.GET("/", root)
	err := http.ListenAndServeTLS(":"+SERVER_PORT, SERVER_CERT, SERVER_KEY, r)
	log.Fatal(err)
}
