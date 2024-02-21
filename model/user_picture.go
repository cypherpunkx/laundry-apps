package model

type UserPicture struct {
	Id string `json:"id"`
	UserId string `json:"userId"`
	FileLocation string `json:"fileLocation"`
}