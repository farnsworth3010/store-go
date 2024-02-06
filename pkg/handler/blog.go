package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"store/models"
	"strconv"
)

type getBlogResponse struct {
	Data  []models.Blog `json:"data"`
	Page  int           `json:"page"`
	Total int           `json:"total"`
}

// @Summary Create blog post
// @Schemes
// @Description creates post
// @Tags blog
// @Param request body models.CreateBlogParams true "query params"
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Router /blog [post]
func (h *Handler) createBlog(c *gin.Context) {
	var input models.CreateBlogParams

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Blog.Create(input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get blog post
// @Schemes
// @Description gets posts
// @Tags blog
// @Param limit   query int true "limit"
// @Param page query int true "page"
// @Produce json
// @Success 200 {object} handler.getBlogResponse
// @Router /blog [get]
func (h *Handler) getBlog(c *gin.Context) {
	page, ok := c.GetQuery("page")
	if !ok {
		NewErrorResponse(c, http.StatusBadRequest, "no page provided")
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "page type conversion error")
	}

	limit, ok := c.GetQuery("limit")
	if !ok {
		NewErrorResponse(c, http.StatusBadRequest, "no offset provided")
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "offset type conversion error")
	}

	blog, total, err := h.services.Blog.Get(pageInt, limitInt)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getBlogResponse{Data: blog, Page: pageInt, Total: total})
}
