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
	router := gin.Default()
	router.Static("/images", "images")
	router.Use(CORSMiddleware())
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
			auth.DELETE("/", h.userIdentity, h.deleteUser)
			auth.GET("/info", h.userIdentity, h.getUserInfo)
		}
		blog := api.Group("/blog")
		{
			blog.POST("/", h.userIdentity, h.createBlog)
			blog.GET("/", h.getBlog)
			//blog.GET("/:id")
			//blog.PUT("/:id")
			blog.DELETE("/:id", h.userIdentity, h.deleteBlog)
		}
		product := api.Group("/product")
		{
			product.POST("/", h.userIdentity, h.createProduct)
			product.POST("/search", h.userIdentity, h.getProductsByName)
			product.GET("/", h.getProduct)
			product.GET("/:id", h.getProductById)
			product.GET("/latest", h.getLatestProduct)
			//product.DELETE("/:id", h.deleteProduct)
		}
		categories := api.Group("/categories")
		{
			categories.POST("/", h.userIdentity, h.addCategory)
			categories.PATCH("/", h.userIdentity, h.updateCategory)
			categories.DELETE("/", h.userIdentity, h.deleteCategory)
			categories.GET("/", h.getCategories)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
