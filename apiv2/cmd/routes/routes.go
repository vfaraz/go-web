package routes

import (
	"rest/cmd/handlers"
	"rest/internal/domain"
	"rest/internal/product"

	"github.com/gin-gonic/gin"
)

type Router struct {
	db     *[]domain.Product
	engine *gin.Engine
}

func NewRouter(en *gin.Engine, db *[]domain.Product) *Router {
	return &Router{engine: en, db: db}
}

func (r *Router) SetRoutes() {
	r.SetProduct()
}

// website
func (r *Router) SetProduct() {
	// instances
	repo := product.NewRepository(r.db, len(*r.db))
	sv := product.NewService(repo)
	handler := handlers.NewProduct(sv)

	p := r.engine.Group("/products")
	p.GET("", handler.GetProducts())
	p.GET(":id", handler.GetProductById())
	p.POST("", handler.CreateProduct())
	p.PUT(":id", handler.UpdateProduct())
	p.PATCH(":id", handler.PartialUpdateProduct())
	p.DELETE(":id", handler.DeleteProduct())

}
