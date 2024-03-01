package handler

import (
	"net/http"
	"store/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type getProductResponse struct {
	Data  []models.Product `json:"data"`
	Page  int              `json:"page"`
	Total int64            `json:"total"`
}

func (h *Handler) createProduct(c *gin.Context) {
	var input models.Product

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Product.Create(input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getProduct(c *gin.Context) {
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

	product, total := h.services.Product.Get(pageInt, limitInt)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getProductResponse{Data: product, Page: pageInt, Total: total})
}

func (h *Handler) getLatestProduct(c *gin.Context) {
	product := h.services.Product.Latest()
	c.JSON(http.StatusOK, getProductResponse{Data: product})
}

func (h *Handler) deleteProduct(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "id type conversion error")
		return
	}
	h.services.Product.Delete(uint(ID))
}

func (h *Handler) getProductById(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "id type conversion error")
		return
	}
	product := h.services.Product.GetById(uint(ID))
	c.JSON(http.StatusOK, product)
}

func (h *Handler) getProductsByName(c *gin.Context) {
	var input models.SearchProductInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	products, err := h.services.Product.GetByName(input.Name)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getProductResponse{Data: products})
}

func (h *Handler) getCategories(c *gin.Context) {
	categories, err := h.services.Product.GetCategories()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (h *Handler) addCategory(c *gin.Context) {
	var input models.CategoryInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	ID, err := h.services.AddCategory(input.Name)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": ID,
	})
}

func (h *Handler) updateCategory(c *gin.Context) {
	var input models.UpdateCategoryInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.UpdateCategory(input.ID, input.Name)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) deleteCategory(c *gin.Context) {
	var input models.DeleteCategoryInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.DeleteCategory(input.ID)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) getBrands(c *gin.Context) {
	brands, err := h.services.Product.GetBrands()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": brands,
	})
}

func (h *Handler) updateProduct(c *gin.Context) {
	var input models.Product

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Product.Update(input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
