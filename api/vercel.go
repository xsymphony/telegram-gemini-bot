package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	gateway *gin.Engine
)

func init() {
	gateway = gin.Default()
	gateway.LoadHTMLGlob("templates/*")
	Register(gateway)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	gateway.ServeHTTP(w, r)
}
