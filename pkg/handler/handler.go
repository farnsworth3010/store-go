package handler

import (
	"github.com/gin-gonic/gin"
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
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	//api := router.Group("/api")
	//{
	//	blog := api.Group("/blog")
	//	{
	//		blog.POST("/")
	//		blog.GET("/")
	//		blog.GET("/:id")
	//		blog.PUT("/:id")
	//		blog.DELETE("/:id")
	//	}
	//}

	return router
}
