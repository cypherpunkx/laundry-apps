package repository

import (
	"database/sql"
	"fmt"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/utils/common"
	"enigmacamp.com/enigma-laundry-apps/utils/constant"
)

type CustomerRepository interface {
	BaseRepository[model.Customer]
	BaseRepositoryPaging[model.Customer]
}

// sedangkan untuk struct ini (object implementasion interface) kita jadikan private,
//
//	sehingga akses tidak sembarangan
type customerRepository struct {
	db *sql.DB
}

func (c *customerRepository) Create(payload model.Customer) error {
	_,err := c.db.Exec(constant.CUSTOMER_INSERT,payload.Id,payload.Name,payload.PhoneNumber,payload.Address)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) List() ([]model.Customer,error){
	return nil,nil
}

func (c *customerRepository) Get(id string) (model.Customer,error) {
	var customer model.Customer
	err := c.db.QueryRow(constant.CUSTOMER_GET,id).Scan(
		&customer.Id,
		&customer.Name,
		&customer.PhoneNumber,
		&customer.Address,
	)
	if err != nil {
		return model.Customer{},fmt.Errorf("Error Get Employee : %s ", err.Error())
	}
	return customer,nil
}

func (c *customerRepository) Update(payload model.Customer) error {
	return nil
}

func (c *customerRepository) Delete(id string) error {
	return nil
}

func (c *customerRepository) Paging(requestPaging dto.PaginationParam, query ...string) ([]model.Customer, dto.Paging, error){
	var paginationQuery dto.PaginationQuery
	paginationQuery = common.GetPaginationParams(requestPaging)
	querySelect := "SELECT id, name, phone_number, address FROM customer LIMIT $1 OFFSET $2"	
	rows, err := c.db.Query(querySelect, paginationQuery.Take, paginationQuery.Skip)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		customers = append(customers, customer)
	}

	// count total rows
	var totalRows int
	row := c.db.QueryRow("SELECT COUNT(*) FROM customer")
	err = row.Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	return customers, common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows), nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

// product 
// customer
// employee

// bill/transaksi