package handler

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"store/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware())
	api := router.Group("/api")
	{
		//auth := api.Group("/auth")
		//{
		//	//auth.POST("/sign-up", h.signUp)
		//	//auth.POST("/sign-in", h.signIn)
		//}
		blog := api.Group("/blog")
		{
			blog.POST("/", h.createBlog)
			blog.GET("/", h.getBlog)
			//blog.GET("/:id")
			//blog.PUT("/:id")
			blog.DELETE("/:id", h.deleteBlog)
		}
		product := api.Group("/product")
		{
			product.POST("/", h.createProduct)
			product.GET("/latest", h.getLatestProduct)
			product.GET("/", h.getProduct)
			//product.DELETE("/:id", h.deleteProduct)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
