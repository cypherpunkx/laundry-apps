package dto

import (
	"time"

	"enigmacamp.com/enigma-laundry-apps/model"
)

type BillResponseDto struct {
	Id       string `json:"id"`
	BillDate time.Time 	`json:"billDate"`
	EntryDate time.Time `json:"entryDate"`
	Employee model.Employee `json:"employee"`
	Customer model.Customer `json:"customer"`
	BillDetails []BillDetailReponseDto `json:"billDetails"`
	TotalBill int `json:"totalBill"`
}

type BillDetailReponseDto struct {
	Id string `json:"id"`
	BillId string `json:"billId"`
	Product	model.Product `json:"product"`
	ProductPrice int `json:"productPrice"`
	Qty int `json:"qty"`
	FinishDate time.Time `json:"finishDate"`
}