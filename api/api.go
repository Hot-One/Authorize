package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageI, logger logger.LoggerI) {

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	handler := handler.NewHandler(cfg, storage, logger)

	// Login Api
	r.POST("/login", handler.Login)

	// Register Api
	r.POST("/register", handler.Register)

	// User Api
	r.POST("/user", handler.AuthMiddleware(), handler.CreateUser)
	r.GET("/user/:id", handler.AuthMiddleware(), handler.GetByIdUser)
	r.GET("/user", handler.GetListUser)
	r.PUT("/user/:id", handler.AuthMiddleware(), handler.UpdateUser)
	r.DELETE("/user/:id", handler.AuthMiddleware(), handler.DeleteUser)

	// Category Api
	r.POST("/category", handler.AuthMiddleware(), handler.CreateCategory)
	r.GET("/category/:id", handler.GetByIdCategory)
	r.GET("/category", handler.GetListCategory)
	r.PUT("/category/:id", handler.AuthMiddleware(), handler.UpdateCategory)
	r.DELETE("/category/:id", handler.AuthMiddleware(), handler.DeleteCategory)

	// Product Api
	r.POST("/product", handler.AuthMiddleware(), handler.CreateProduct)
	r.GET("/product/:id", handler.GetByIdProduct)
	r.GET("/product", handler.GetListProduct)
	r.PUT("/product/:id", handler.AuthMiddleware(), handler.UpdateProduct)
	r.PATCH("/product/:id", handler.AuthMiddleware(), handler.PatchProduct)
	r.DELETE("/product/:id", handler.AuthMiddleware(), handler.DeleteProduct)

	// Market Api
	r.POST("/market", handler.AuthMiddleware(), handler.CreateMarket)
	r.GET("/market/:id", handler.AuthMiddleware(), handler.GetByIdMarket)
	r.GET("/market", handler.AuthMiddleware(), handler.GetListMarket)
	r.PUT("/market/:id", handler.AuthMiddleware(), handler.UpdateMarket)
	r.PATCH("/market/:id", handler.AuthMiddleware(), handler.PatchMarket)
	r.DELETE("/market/:id", handler.AuthMiddleware(), handler.DeleteMarket)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Accesp-Encoding, Authorization, Cache-Control")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
