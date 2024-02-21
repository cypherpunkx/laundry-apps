package repository

import (
	"mime/multipart"
	"path/filepath"

	"enigmacamp.com/enigma-laundry-apps/utils/common"
	"github.com/gin-gonic/gin"
)

type FileRepository interface {
	Save(fileName string, file *multipart.File) (string, error)
	FindFile(c *gin.Context,filepath string, filename string) 
}

type fileRepository struct {
	fileBasePath string
}

func (f *fileRepository) Save(fileName string, file *multipart.File) (string, error) {
	fileLocation := filepath.Join(f.fileBasePath, fileName)
	err := common.SaveToLocalFile(fileLocation, file)
	if err != nil {
		return "", err
	}
	return fileLocation, nil
}

func (f *fileRepository) FindFile(c *gin.Context,filepath string, filename string) {
	c.FileAttachment(filepath,filename)
}

func NewFileRepository(basePath string) FileRepository {
	fileRepo := fileRepository{fileBasePath: basePath}
	return &fileRepo
}
