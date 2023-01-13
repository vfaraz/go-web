package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"rest/cmd/handlers"
	"rest/internal/domain"
	"rest/internal/product"
	"rest/pkg/store"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type response struct {
	Data interface{} `json:"data"`
}

func createServer(token string) *gin.Engine {

	if token != "" {
		err := os.Setenv("TOKEN", token)
		if err != nil {
			panic(err)
		}
	}

	db := store.NewStorage("./products_test.json")
	repo := product.NewRepository(db)
	service := product.NewService(repo)
	productHandler := handlers.NewProduct(service)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	pr := r.Group("/products")
	{
		pr.GET("", productHandler.GetProducts())
		//pr.GET(":id", productHandler.GetByID())
		//pr.GET("/search", productHandler.Search())
		//pr.POST("", productHandler.Post())
		//pr.DELETE(":id", productHandler.Delete())
		//pr.PATCH(":id", productHandler.Patch())
		//pr.PUT(":id", productHandler.Put())
	}
	return r
}

func createRequestTest(method string, url string, body string, token string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	if token != "" {
		req.Header.Add("TOKEN", token)
	}
	return req, httptest.NewRecorder()
}

func loadProducts(path string) ([]domain.Product, error) {
	var products []domain.Product
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func Test_GetAll_OK(t *testing.T) {
	var expectd = response{Data: []domain.Product{}}

	r := createServer("my-secret-token")
	request, rr := createRequestTest(http.MethodGet, "/products", "", "my-secret-token")

	products, err := loadProducts("./products_copy.json")
	if err != nil {
		panic(err)
	}
	expectd.Data = products
	actual := map[string][]domain.Product{}

	r.ServeHTTP(rr, request)

	assert.Equal(t, 200, rr.Code)
	err = json.Unmarshal(rr.Body.Bytes(), &actual)
	assert.Nil(t, err)
	assert.Equal(t, expectd.Data, actual["data"])
}
