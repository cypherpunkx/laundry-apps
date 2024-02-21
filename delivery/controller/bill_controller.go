package controller

import (
	"net/http"
	"time"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/usecase"
	"enigmacamp.com/enigma-laundry-apps/utils/common"
	"github.com/gin-gonic/gin"
)

type BillController struct {
	router *gin.Engine
	billUc usecase.BillUseCase
}

func (b *BillController) createHandler(c *gin.Context) {
	var bill model.Bill
	err := c.ShouldBindJSON(&bill)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err" : err.Error(),
		})
		return
	}
	bill.Id = common.GenerateUUID()
	if err := b.billUc.RegisterNewBill(bill); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err" : err.Error(),
		})
		return
	}

	type BillDetails struct {
		ProductId string `json:"productId"`
		FinishDate time.Time `json:"finishDate"`
		Qty int `json:"qty"`
	}

	var billDetails []BillDetails
	for _, item := range bill.BillDetails {
		billDetails = append(billDetails, BillDetails{
			ProductId: item.ProductId,
			FinishDate: item.FinishDate,
			Qty: item.Qty,
		})
	}

	billResponse := map[string]any {
		"id" : bill.Id,		
		"employeeId" : bill.EmployeeId,
		"customerId" : bill.CustomerId,
		"billdetails" : billDetails,
	}
	c.JSON(http.StatusOK,billResponse)
}

func (b *BillController) getHandler(c *gin.Context) {
	billId := c.Param("id")
	billResponseDto,err := b.billUc.FindBillById(billId)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message" : "Success Get Bill",
		"data" : billResponseDto,
	})
}

func NewBillController(r *gin.Engine, usecase usecase.BillUseCase) {
	controller := BillController{
		router: r,
		billUc: usecase,
	}
	rg := controller.router.Group("/api/v1")
	rg.POST("/bills",controller.createHandler)
	rg.GET("/bills/:id",controller.getHandler)
}