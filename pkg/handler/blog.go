package handler

import (
	"net/http"
	"store/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type getBlogResponse struct {
	Data  []models.Blog `json:"data"`
	Page  int           `json:"page"`
	Total int64         `json:"total"`
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
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "page type conversion error")
		return
	}

	limit, ok := c.GetQuery("limit")
	if !ok {
		NewErrorResponse(c, http.StatusBadRequest, "no offset provided")
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "offset type conversion error")
		return
	}

	blog, total := h.services.Blog.Get(pageInt, limitInt)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getBlogResponse{Data: blog, Page: pageInt, Total: total})
}

func (h *Handler) deleteBlog(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "id type conversion error")
	}
	h.services.Blog.Delete(uint(ID))
}

func (h *Handler) updateBlog(c *gin.Context) {
	var input models.EditBlogParams

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Blog.Update(input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}
