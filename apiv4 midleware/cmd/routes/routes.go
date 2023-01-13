package routes

import (
	"os"
	"rest/cmd/handlers"
	"rest/cmd/middlewares"

	"rest/internal/product"
	"rest/pkg/store"

	"github.com/gin-gonic/gin"

	"rest/cmd/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	//url := ginSwagger.URL("localhost:8080") // The url pointing to API definition
	r.engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	p := r.engine.Group("/products")
	p.Use(middlewares.TokenAuthMiddleware())
	{
		p.GET("", handler.GetProducts(), middlewares.LoggerMiddleware())
		p.GET(":id", handler.GetProductById())
		p.POST("", handler.CreateProduct())
		p.PUT(":id", handler.UpdateProduct())
		p.PATCH(":id", handler.PartialUpdateProduct())
		p.DELETE(":id", handler.DeleteProduct())
	}

}
