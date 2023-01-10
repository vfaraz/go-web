package handlers

import (
	"fmt"
	"net/http"
	"rest/pkg/response"
	"rest/services"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Name        string  `json:"name" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	CodeValue   int     `json:"code_value" validate:"required"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

func CreateProduct(c *gin.Context) {
	// request
	var body Request
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err))
		return
	}
	validate := validator.New()
	if err := validate.Struct(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Error(err))
		return
	}
	date := strings.Split(body.Expiration, "/")
	_, err := time.Parse("2006-01-02", fmt.Sprintf("%s-%s-%s", date[2], date[1], date[0]))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err))
		return
	}
	// process
	prod, err := services.Create(body.Name, body.Quantity, body.CodeValue,
		body.IsPublished, body.Expiration, body.Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	// response
	c.JSON(http.StatusCreated,
		response.Ok("succed create product", prod))

}

func GetProductById(c *gin.Context) {
	// request
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err))
		return
	}
	// process
	product, err := services.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err))
		return
	}
	// response
	c.JSON(http.StatusOK, response.Ok("succed to get product", product))

}

func GetProducts(c *gin.Context) {
	// request

	// process
	products := services.Get()
	// response
	c.JSON(http.StatusOK, response.Ok("succed to get products", products))
}
