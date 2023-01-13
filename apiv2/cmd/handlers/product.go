package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"rest/internal/product"
	"rest/pkg/response"
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

type Product struct {
	sv product.Service
}

func NewProduct(sv product.Service) *Product {
	return &Product{sv: sv}
}
func (p *Product) GetProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request

		// process
		product, err := p.sv.Get()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.Error(err))
			return
		}

		// response
		ctx.JSON(http.StatusOK, response.Ok("succed to get products", product))

	}
}

func (p *Product) CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {

		// request
		var body Request
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, response.Error(err))
			return
		}
		// validator
		validate := validator.New()
		if err := validate.Struct(&body); err != nil {
			c.JSON(http.StatusUnprocessableEntity, response.Error(err))
			return
		}
		//check date
		date := strings.Split(body.Expiration, "/")
		_, err := time.Parse("2006-01-02", fmt.Sprintf("%s-%s-%s", date[2], date[1], date[0]))
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error(err))
			return
		}
		// process
		prod, err := p.sv.Create(body.Name, body.Quantity, body.CodeValue,
			body.IsPublished, body.Expiration, body.Price)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error(err))
			return
		}

		// response
		c.JSON(http.StatusCreated,
			response.Ok("success to create product", prod))

	}
}

func (p *Product) GetProductById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request
		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error(err))
			return
		}
		// process
		product, err := p.sv.GetByID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error(err))
			return
		}
		// response
		c.JSON(http.StatusOK, response.Ok("success to get product", product))

	}
}

func (p *Product) UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request
		var body Request
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, response.Error(err))
			return
		}
		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error(err))
			return
		}
		// process
		product, err := p.sv.Update(id, body.Name, body.Quantity, body.CodeValue,
			body.IsPublished, body.Expiration, body.Price)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error(err))
			return
		}
		// response
		c.JSON(http.StatusOK, response.Ok("success to update product", product))

	}
}

func (p *Product) PartialUpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request
		// obtenemos el id
		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error(err))
			return
		}
		// process
		// obtenemos el producto sin actualizar
		product, err := p.sv.GetByID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error(err))
			return
		}
		// tomando el producto modificamos solo los campos que nos envian
		err = json.NewDecoder(c.Request.Body).Decode(&product)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error(err))
			return
		}
		product, err = p.sv.PartialUpdate(product)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error(err))
			return
		}
		// response
		c.JSON(http.StatusOK, response.Ok("success to update product", product))

	}
}

func (p *Product) DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request
		// obtenemos el id
		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error(err))
			return
		}
		// process
		err = p.sv.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error(err))
			return
		}
		// response
		c.JSON(http.StatusOK, response.Ok("success to delete product", nil))

	}

}
