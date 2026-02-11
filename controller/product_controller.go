package controller

import (
	"GoApi/model"
	"net/http"

	"github.com/gin-gonic/gin"
)	

type ProductController struct {
	//Usecase
}

func NewProductController() ProductController {
	return ProductController{}
}

func (p *ProductController) GetProducts(ctx *gin.Context){

	products := []model.Product{
		{ID: 1, Name: "Product 1", Price: 10.99},
		{ID: 2, Name: "Product 2", Price: 19.99},
	}

	ctx.JSON(http.StatusOK, products)
}