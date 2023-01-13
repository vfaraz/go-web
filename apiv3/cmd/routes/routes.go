package routes

import (
	"rest/cmd/handlers"
	//"rest/internal/domain"
	"rest/internal/product"
	"rest/pkg/store"

	"github.com/gin-gonic/gin"
)

type Router struct {
	storage store.Storage
	engine  *gin.Engine
}

func NewRouter(en *gin.Engine, storage store.Storage) *Router {
	return &Router{engine: en, storage: storage}
}

func (r *Router) SetRoutes() {
	r.SetProduct()
}

// website
func (r *Router) SetProduct() {
	// instances
	repo := product.NewRepository(r.storage)
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
