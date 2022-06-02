package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float64 `json:"price"`
	Code         string  `json:"code"`
	Published    bool    `json:"published"`
	CreationDate string  `json:"creationDate"`
}

func readProductFile() (data []byte, err error) {
	data, err = ioutil.ReadFile("./bootcamp/goweb1/products.json")

	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetAll(c *gin.Context) {
	data, err := readProductFile()
	var products []product

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro":        err.Error(),
			"mensagem":    "Ocorreu um erro ao trazer os produtos",
			"status_code": http.StatusInternalServerError,
		})
		return
	}
	json.Unmarshal(data, &products)
	c.JSON(http.StatusOK, products)
	return
}

func GetOne(c *gin.Context) {
	data, err := readProductFile()

	var products []product

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro":        err.Error(),
			"mensagem":    "Ocorreu um erro ao trazer os produtos",
			"status_code": http.StatusInternalServerError,
		})
		return
	} else {
		json.Unmarshal(data, &products)
	}

	productId, err2 := strconv.Atoi(c.Param("id"))

	if err2 == nil {
		for _, p := range products {
			if p.Id == productId {
				c.JSON(http.StatusOK, p)
				return
			}
		}
	}
}

func main() {

	router := gin.Default()

	group := router.Group("/products")
	{
		group.GET("/", GetAll)
		group.GET("/:id", GetOne)
	}

	router.Run()
}
