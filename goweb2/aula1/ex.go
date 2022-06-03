package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Quantity int     `json:"quantity" binding:"required"`
}

var products = []product{}

func saveProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		var req product

		if err := c.ShouldBindJSON(&req); err != nil {

			validationString := "Error:Field validation"
			errorMessage := err.Error()

			if strings.Contains(errorMessage, validationString) {
				//acha o indice inicial da string
				errorIndex := strings.Index(errorMessage, validationString)

				// recortar a string para recuperar o nome do campo
				field := errorMessage[14 : errorIndex-2]

				c.JSON(http.StatusBadRequest, gin.H{
					"error": fmt.Sprintf("O campo %s é obrigatório", strings.ToLower(field)),
				})
				return
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"error": errorMessage,
			})
			return
		}

		req.ID = len(products) + 1
		products = append(products, req)
		c.JSON(http.StatusOK, req)
	}
}

func getAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"produtos": &products,
	})
}

func main() {
	router := gin.Default()

	routesGroup := router.Group("/products")
	{
		routesGroup.GET("/", getAllProducts)
		routesGroup.POST("/", saveProducts())

	}

	router.Run()
}
