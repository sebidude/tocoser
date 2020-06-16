package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

var (
	contentDir string
	token      string
	acmeDir    string
	dnsname    string
)

func GetEnvOrDefault(key, defaultValue string) string {
	v := os.Getenv(key)
	if len(v) > 0 {
		return v
	}
	return defaultValue
}

func main() {
	contentDir = GetEnvOrDefault("CONTENTDIR", "content")
	token = GetEnvOrDefault("TOKEN", "1234")
	acmeDir = os.Getenv("ACMEDIR")
	dnsname = os.Getenv("DNSNAME")
	authuser := os.Getenv("AUTHUSER")
	authpassword := os.Getenv("AUTHPASSWORD")

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	if len(authuser) != 0 && len(authpassword) != 0 {

		authorized := r.Group("/auth/content", gin.BasicAuth(gin.Accounts{
			authuser: authpassword,
		}))

		authorized.StaticFS("", gin.Dir(contentDir, true))
	}

	tocoser := r.Group("/token/content", checkToken)
	tocoser.StaticFS("", gin.Dir(contentDir, true))

	if len(dnsname) == 0 || len(acmeDir) == 0 {
		log.Fatalln(r.Run(GetEnvOrDefault("LISTEN_ADDRESS", ":8080")))
	} else {
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(dnsname),
			Cache:      autocert.DirCache(acmeDir),
		}
		log.Fatal(autotls.RunWithManager(r, &m))
	}

}

func checkToken(c *gin.Context) {
	authToken := c.GetHeader("x-auth-token")
	if len(authToken) == 0 || authToken != token {
		c.String(http.StatusUnauthorized, "Go away!\n")
		c.Abort()
	}
}
