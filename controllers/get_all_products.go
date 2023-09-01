package controllers

import (
	"net/http"

	"github.com/Rodrigosiq03/serverless-go-api/entities"
	"github.com/gin-gonic/gin"
)

func GetAllProducts(context *gin.Context) {
	var products []entities.Product

	products = append(products, entities.Product{
		Name:        "Product 1",
		Description: "Description 1",
		Price:       9.99,
	})

	context.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}
