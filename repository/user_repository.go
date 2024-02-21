package repository

import (
	"database/sql"
	"fmt"
	"enigmacamp.com/enigma-laundry-apps/model"	
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(payload model.UserCredential) error
	List() ([]model.UserCredential,error)
	GetByUsername(username string) (model.UserCredential,error)
	GetByUsernamePassword(username string, password string) (model.UserCredential,error)	
	GetById(id string) (model.UserCredential,error)
}


type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Create(payload model.UserCredential) error {	
	_,err := u.db.Exec("INSERT INTO user_credential (id,username,password) values ($1,$2,$3)",payload.Id,payload.Username,payload.Password)
	if err != nil {
		return nil
	}
	return nil
}

func (u *userRepository) List() ([]model.UserCredential,error) {
	var users []model.UserCredential
	rows, err := u.db.Query("SELECT id,username, is_active from user_credential")
	if err != nil {
		return nil,err
	}
	for rows.Next() {
		var user model.UserCredential
		err := rows.Scan(&user.Id,&user.Username,&user.IsActive)
		if err != nil {
			return nil,fmt.Errorf("Error scan user : %s", err.Error())
		}
		users = append(users,user)
	}
	return users,nil
}



func (u *userRepository) GetByUsername(username string) (model.UserCredential,error) {
	var user model.UserCredential
	err := u.db.QueryRow("SELECT id,username,password,is_active FROM user_credential WHERE is_active = $1 AND username = $2",true,username).Scan(&user.Id,&user.Username,&user.Password,&user.IsActive)
	if err != nil {
		return model.UserCredential{},err
	}
	return user,nil
}

func (u *userRepository) GetByUsernamePassword(username string, password string) (model.UserCredential,error) {
	user, err := u.GetByUsername(username)
	if err != nil {
		return model.UserCredential{},err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.UserCredential{},err
	}
	return user,nil
}

func (u *userRepository) GetById(id string) (model.UserCredential,error) {
	var userCred model.UserCredential
	row := u.db.QueryRow("SELECT id,username,is_active FROM user_credential where id=$1",id)
	err := row.Scan(&userCred.Id,&userCred.Username,&userCred.IsActive)
	if err != nil {
		return model.UserCredential{},err
	}
	return userCred,nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}