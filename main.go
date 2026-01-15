package main

import (
	"crypto/elliptic"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/html"
	"gopkg.in/yaml.v2"
)

func main() {
	// Using crypto/elliptic (vulnerable to GO-2022-0435)
	curve := elliptic.P256()
	fmt.Printf("Using elliptic curve: %v\n", curve.Params().Name)

	// Using crypto/tls (vulnerable to GO-2022-0434)
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	fmt.Printf("TLS config min version: %d\n", tlsConfig.MinVersion)

	// Using net/http/httputil ReverseProxy (vulnerable to GO-2022-0520)
	target, _ := url.Parse("http://example.com")
	proxy := httputil.NewSingleHostReverseProxy(target)
	fmt.Printf("Created reverse proxy to: %s\n", target.String())

	// Using net/http client (vulnerable to GO-2022-0434)
	client := &http.Client{}
	fmt.Printf("HTTP client created: %v\n", client)

	// Using vulnerable external dependencies
	_ = gin.Default()
	_ = websocket.Upgrader{}
	_, _ = bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	_ = html.NewTokenizer(nil)

	data := struct{ Name string }{Name: "test"}
	_, _ = yaml.Marshal(&data)

	fmt.Println("This project uses Go 1.17 stdlib which has known vulnerabilities")
	fmt.Println("Dependabot should detect: GO-2022-0434, GO-2022-0435, GO-2022-0520, etc.")
	fmt.Println("Also includes vulnerable external dependencies from 2019-2021")
}
