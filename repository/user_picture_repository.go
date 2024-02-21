package repository

import (
	"database/sql"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/utils/constant"
)

type UserPicture interface {
	Create(userPicture model.UserPicture) error
	Get(id string) (model.UserPicture,error)
}

type userPicture struct {
	db *sql.DB
}

func (u *userPicture) Create(userPicture model.UserPicture) error {
	_,err := u.db.Exec(constant.CREATE_USER_PICTURE,userPicture.Id,userPicture.UserId,userPicture.FileLocation)
	if err != nil {
		return err
	}
	return nil
}

func (u *userPicture)Get(id string) (model.UserPicture,error) {
	var userPicure model.UserPicture
	row := u.db.QueryRow(constant.GET_USER_PICTURE,id)
	err := row.Scan(&userPicure.Id,&userPicure.UserId,&userPicure.FileLocation)
	if err != nil {
		return model.UserPicture{},err
	}
	return userPicure,nil
}

func NewUserPicrerRepository(db *sql.DB) UserPicture {
	return &userPicture {
		db: db,
	}
}