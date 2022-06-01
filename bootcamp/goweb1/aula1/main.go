package main

import (
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

	//Capture o request GET "/ola-leandro"
	router.GET("/ola-leandro", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá Leandro!",
		})
	})

	router.GET("/products", func(getAll *gin.Context) {
		getAll.File("./bootcamp/goweb1/products.json")
	})

	// Executamos nosso servidor na porta 8080
	router.Run()
}
