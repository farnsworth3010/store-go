package handler

import (
	"store/pkg/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
			blog.PATCH("/", h.userIdentity, h.updateBlog)
			blog.DELETE("/:id", h.userIdentity, h.deleteBlog)
		}
		brand := api.Group("/brand")
		{
			brand.POST("/", h.userIdentity, h.createBrand)
			brand.GET("/", h.getBrand)
			brand.PATCH("/", h.userIdentity, h.updateBrand)
			brand.DELETE("/:id", h.userIdentity, h.deleteBrand)
		}
		product := api.Group("/product")
		{
			product.GET("/brands", h.getBrands)
			product.POST("/filter", h.filterProducts)
			product.POST("/", h.userIdentity, h.createProduct)
			product.PATCH("/", h.userIdentity, h.updateProduct)
			product.POST("/search", h.userIdentity, h.getProductsByName)
			product.GET("/", h.getProduct)
			product.GET("/:id", h.getProductById)
			product.GET("/latest", h.getLatestProduct)
			product.DELETE("/:id", h.userIdentity, h.deleteProduct)
		}
		categories := api.Group("/categories")
		{
			categories.POST("/", h.userIdentity, h.addCategory)
			categories.PATCH("/", h.userIdentity, h.updateCategory)
			categories.DELETE("/", h.userIdentity, h.deleteCategory)
			categories.GET("/", h.getCategories)
		}
		admin := api.Group("/panel")
		{
			admin.GET("/admins", h.userIdentity, h.getAdmins)
			admin.GET("/users", h.userIdentity, h.getUsers)
			admin.GET("/blogs", h.userIdentity, h.getBlogs)
			admin.POST("/setRole", h.userIdentity, h.setRole)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
