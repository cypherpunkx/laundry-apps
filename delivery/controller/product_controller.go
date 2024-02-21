package controller

import (
	"net/http"
	"strconv"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/usecase"
	"enigmacamp.com/enigma-laundry-apps/utils/common"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	router *gin.Engine
	productUc usecase.ProductUseCase
}

func (p *ProductController) createHandler(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	product.Id = common.GenerateUUID()
	err := p.productUc.RegisterNewProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status" : http.StatusCreated,
		"message" : "Success Create New Product",
		"data" : product,
	})
}

func (p *ProductController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))	
	paginationParam := dto.PaginationParam{
		Page : page,
		Limit : limit,
	}
	products,paging, err := p.productUc.FindAllProduct(paginationParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get All Data Successfully",
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data":   products,
		"paging": paging,
	})
}

func NewProductController( router *gin.Engine,productUseCase usecase.ProductUseCase) {
	ctr := &ProductController{
		router:  router,
		productUc: productUseCase,
	}

	routerGroup := ctr.router.Group("/api/v1")
	routerGroup.POST("/product",ctr.createHandler)
	routerGroup.GET("/product",ctr.listHandler)
}

