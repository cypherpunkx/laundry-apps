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

type CustomerController struct {
	router *gin.Engine
	customerUc usecase.CustomerUseCase
}

func (cc *CustomerController) createHandler(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	customer.Id = common.GenerateUUID()
	err := cc.customerUc.RegisterNewCustomer(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status" : http.StatusCreated,
		"message" : "Success Create New customer",
		"data" : customer,
	})
}

func (cc *CustomerController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))	
	paginationParam := dto.PaginationParam{
		Page : page,
		Limit : limit,
	}
	customers,paging, err := cc.customerUc.FindAllCustomer(paginationParam)
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
		"data":   customers,
		"paging": paging,
	})
}

func NewCustomerController( router *gin.Engine,customerUseCase usecase.CustomerUseCase) {
	ctr := &CustomerController{
		router:  router,
		customerUc: customerUseCase,
	}

	routerGroup := ctr.router.Group("/api/v1")
	routerGroup.POST("/customer",ctr.createHandler)
	routerGroup.GET("/customer",ctr.listHandler)
}

