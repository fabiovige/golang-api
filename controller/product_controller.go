package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	newProduct, err := p.productUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, newProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do produto deve ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUsecase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do produto deve ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var product model.Product
	err = ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	product.ID = productId

	productObj, err := p.productUsecase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if productObj == nil {
		response := model.Response{
			Message: "Produto não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	updatedProduct, err := p.productUsecase.UpdateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}

func (p *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do produto deve ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productObj, err := p.productUsecase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if productObj == nil {
		response := model.Response{
			Message: "Produto não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	err = p.productUsecase.DeleteProduct(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	response := model.Response{
		Message: "Produto deletado com sucesso",
	}

	ctx.JSON(http.StatusOK, response)
}
