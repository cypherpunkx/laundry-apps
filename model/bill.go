package model

import "time"

type Bill struct {
	Id       string `json:"id"`
	BillDate time.Time `json:"billDate"`
	EntryDate time.Time `json:"entryDate"`	
	EmployeeId string `json:"employeeId"`
	CustomerId string `json:"customerId"`	
	BillDetails []BillDetail `json:"billDetails"`
}

type BillDetail struct {
	Id string `json:"id"`
	BillId string `json:"billId"`
	ProductId string `json:"productId"`
	ProductPrice int `json:"productPrice"`
	Qty int `json:"qty"`
	FinishDate time.Time `json:"finishDate"`
}

// {
// 	"baju" -> 2 kg -> reguler 3 hari =  billDetail 7
// 	"jaket" -> pcs -> expres 1 hari = billDetail 5
// }


// dto ->  Data Transfer Object 
// struct yang akan dijadikan type data untuk request atau response


