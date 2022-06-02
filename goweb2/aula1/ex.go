package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

var products = []product{}

func saveProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req product

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		req.ID = len(products) + 1
		products = append(products, req)
		fmt.Println(products)
		c.JSON(http.StatusOK, products)
	}
}

func main() {
	router := gin.Default()

	routesGroup := router.Group("/products")
	{
		routesGroup.POST("/", saveProducts())

	}

	router.Run()
}
