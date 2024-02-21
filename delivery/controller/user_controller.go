package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/usecase"
	"enigmacamp.com/enigma-laundry-apps/utils/common"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	router *gin.Engine
	userUc usecase.UserUseCase
	userPic usecase.UserPictureUseCase
}

func (u *UserController) createHandler(c *gin.Context) {
	var user model.UserCredential
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error" : err.Error(),
		})
		return
	}

	user.Id = common.GenerateUUID()
	findUser,err := u.userUc.RegisterNewUser(user);
	if  err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error" : err.Error(),
		})
		return
	}

	userResponse := map[string]any {
		"id" : findUser.Id,
		"username" : findUser.Username,
		"isActive" : findUser.IsActive,
	}
	c.JSON(http.StatusOK,userResponse)
}

func (u *UserController) UploadPictureHandler(c *gin.Context) {
	var userPicture model.UserPicture
	id := c.Param("id")
	file, fileHeader, err := c.Request.FormFile("photo")
	if err != nil {		
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	allowedExts := []string{".png", ".jpg", ".jpeg"}
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	allowed := false
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			allowed = true
			break
		}
	}
	if !allowed {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Extention is not allowed"})
		return
	}
	userPicture.Id = common.GenerateUUID()
	userPicture.UserId = id
	err = u.userPic.UploadUserPicture(userPicture,&file,ext)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "File Uploaded",			
	})
}


func (u *UserController) ListHandler(c *gin.Context) {
	users, err := u.userUc.FindAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status" : http.StatusOK,
		"message" : "Success Get List Users",
		"data" : users,
	})
}


func (u *UserController) downloadPictureHandler(c *gin.Context) {
	id := c.Param("id")
	userPic, err := u.userPic.FindUserPictureById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error" : err.Error(),
		})
		return
	}
	fmt.Println("Test : ",userPic)
	c.FileAttachment(userPic.FileLocation,filepath.Base(userPic.FileLocation) )
}


func NewUserController(r *gin.Engine,userUseCase usecase.UserUseCase, userPicture usecase.UserPictureUseCase) {
	controller := UserController {
		router: r,
		userUc: userUseCase,
		userPic: userPicture,
	}
	rg := r.Group("/api/v1")
	rg.GET("/users",controller.ListHandler)
	rg.POST("/users", controller.createHandler)
	rg.POST("/users/upload-photo/:id",controller.UploadPictureHandler)
	rg.GET("/users/download-photo/:id",controller.downloadPictureHandler)
}