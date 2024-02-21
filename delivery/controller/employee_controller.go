package controller

import (
	"net/http"
	"strconv"

	"enigmacamp.com/enigma-laundry-apps/delivery/middleware"
	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/usecase"
	"enigmacamp.com/enigma-laundry-apps/utils/common"
	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	router *gin.Engine
	useCase usecase.EmployeeUseCase
}

func (e *EmployeeController) createHandler(c *gin.Context){
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	employee.Id = common.GenerateUUID()
	err := e.useCase.RegisterNewEmployee(employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status" : http.StatusCreated,
		"message" : "Success Create New Employee",
		"data" : employee,
	})
}

func (e *EmployeeController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	name := c.Query("name")
	paginationParam := dto.PaginationParam{
		Page : page,
		Limit : limit,
	}
	employees,paging, err := e.useCase.FindAllEmployee(paginationParam,name)
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
		"data":   employees,
		"paging": paging,
	})
}

func (e *EmployeeController) getHandler(c *gin.Context) {
	id := c.Param("id")
	employee, err := e.useCase.FindEmployeeById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status" : http.StatusOK,
		"message" : "Success Get Employee by Id",
		"data" : employee,
	})
}

func (e *EmployeeController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	 err := e.useCase.DeleteEmployee(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status" : http.StatusOK,
		"message" : "Success Delete",		
	})
}

func (e *EmployeeController) updateHandler(c *gin.Context){
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}	
	err := e.useCase.UpdateEmployee(employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status" : http.StatusCreated,
		"message" : "Success Updated Employee",
		"data" : employee,
	})
}


func NewEmployeeController( router *gin.Engine,emplUseCase usecase.EmployeeUseCase) {
	ctr := &EmployeeController{
		router:  router,
		useCase: emplUseCase,
	}

	routerGroup := ctr.router.Group("/api/v1",middleware.AuthMiddleware())
	routerGroup.POST("/employee",ctr.createHandler)
	routerGroup.GET("/employee",ctr.listHandler)
	routerGroup.GET("/employee/:id",ctr.getHandler)
	routerGroup.PUT("/employee",ctr.updateHandler)
	routerGroup.DELETE("/employee/:id",ctr.deleteHandler)
}