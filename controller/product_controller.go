package controller

import (
	"GoApi/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)	

type ProductController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(productUseCase usecase.ProductUseCase) ProductController {
	return ProductController{
		productUseCase: productUseCase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context){
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
		return
	}
	ctx.JSON(http.StatusOK, products)
}