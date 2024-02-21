package usecase

import (
	"fmt"
	"mime/multipart"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/repository"
)

type UserPictureUseCase interface {
	UploadUserPicture(userPicture model.UserPicture,file *multipart.File, extFile string) error
	FindUserPictureById(id string) (model.UserPicture,error)
}

type userPictureUseCase struct {
	repo repository.UserPicture
	fileRepo repository.FileRepository
	userUC UserUseCase
}

func (u *userPictureUseCase)UploadUserPicture(userPicture model.UserPicture,file *multipart.File, extFile string) error{
	userCred, err := u.userUC.FindById(userPicture.UserId)
	if err != nil {
		return err
	}
	fileName := fmt.Sprintf("%s-%s%s",userCred.Username,userPicture.UserId,extFile)	
	filePath, err := u.fileRepo.Save(fileName,file)
	if err != nil {
		return err
	}
	userPicture.FileLocation = filePath
	err = u.repo.Create(userPicture)
	if err != nil {
		return fmt.Errorf("Failed Upload : %s",err.Error())
	}
	return nil
}


func (u *userPictureUseCase)FindUserPictureById(id string) (model.UserPicture,error) {
	userPicture,err := u.repo.Get(id)
	if err != nil {
		return model.UserPicture{}, fmt.Errorf("Failed Get Picture By Id : %s",err.Error())
	}
	return userPicture,nil
}


func NewUserPictureUseCase(repository repository.UserPicture, fileRepository repository.FileRepository, userUseCase UserUseCase) UserPictureUseCase {
	return &userPictureUseCase{
		repo: repository,
		fileRepo: fileRepository,
		userUC: userUseCase,
	}
}