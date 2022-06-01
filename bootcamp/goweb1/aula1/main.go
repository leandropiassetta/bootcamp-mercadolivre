package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// net/http -> permite gerar servidores web
// Um conceito fundamental em servidores net/http são os handlers

type product struct {
	Id           int
	Name         string
	Color        string
	Price        float64
	Code         string
	Published    bool
	CreationDate string
}

func main() {
	router := gin.Default()

	router.GET("/ola-leandro", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá Leandro!",
		})
	})

	http.ListenAndServe(":8080", nil)
}
