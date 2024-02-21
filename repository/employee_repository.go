package repository

import (
	"database/sql"
	"fmt"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/utils/common"
	"enigmacamp.com/enigma-laundry-apps/utils/constant"
)

// interface ini merupakan Kontrak semua method employee repository, sehingga untuk komunikasi antar layer
// menggunakan tipe data interface ini ya, ini alasannya mengapa interface kita dibuat public
type EmployeeRepository interface {
	BaseRepository[model.Employee]
	BaseRepositoryPaging[model.Employee]
}

// sedangkan untuk struct ini (object implementasion interface) kita jadikan private,
//  sehingga akses tidak sembarangan
type employeeRepository struct {
	db *sql.DB
}

func (e *employeeRepository) Create(payload model.Employee) error{
	_,err := e.db.Exec(constant.EMPLOYEE_INSERT,payload.Id,payload.Name,payload.PhoneNumber,payload.Address)
	if err != nil {
		return err
	}
	return nil
}

func(e *employeeRepository) Paging(requestPaging dto.PaginationParam, query ...string) ([]model.Employee, dto.Paging, error){
	var paginationQuery dto.PaginationQuery
	paginationQuery = common.GetPaginationParams(requestPaging)
	querySelect := "SELECT id, name, phone_number, address FROM employee"
	if query[0] != "" {	
		querySelect += ` WHERE name ilike '%` + query[0] + `%'`
	}
	querySelect += ` LIMIT $1 OFFSET $2`
	rows, err := e.db.Query(querySelect, paginationQuery.Take, paginationQuery.Skip)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
		err := rows.Scan(&employee.Id, &employee.Name, &employee.PhoneNumber, &employee.Address)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		employees = append(employees, employee)
	}

	// count total rows
	var totalRows int
	row := e.db.QueryRow("SELECT COUNT(*) FROM employee")
	err = row.Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	return employees, common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows), nil
}

func (e *employeeRepository) List() ([]model.Employee, error){	
	rows, err := e.db.Query(constant.EMPLOYEE_LIST)
	if err != nil {
		return nil,err
	}
	var employees []model.Employee
	
	for rows.Next() {
		var employee model.Employee		
		err = rows.Scan(&employee.Id,&employee.Name, &employee.PhoneNumber, &employee.Address)
		employees = append(employees,employee)
	}
	return employees, nil
}

func(e *employeeRepository) Get(id string) (model.Employee,error){
	// routernya ex : /api/v1/employee/:id
	// localhost:8080/api/v1/employee/s090-2344-dodif-9898
	// (/s090-2344-dodif-9898) diambil menggunakan params dari *gin.Context 
	// ex : .Param("id")
	var employee model.Employee
	err := e.db.QueryRow(constant.EMPLOYEE_GET,id).Scan(
		&employee.Id,
		&employee.Name,
		&employee.PhoneNumber,
		&employee.Address,
	)
	if err != nil {
		return model.Employee{},fmt.Errorf("Error Get Employee : %s ", err.Error())
	}
	return employee,nil
}

func(e *employeeRepository) Update(payload model.Employee) error {
	_,err := e.db.Exec(constant.EMPLOYEE_UPDATE,payload.Name,payload.PhoneNumber,payload.Address, payload.Id)
	if err != nil {
		return fmt.Errorf(" Error Update Employee : %s ", err.Error())
	}
	return nil
}

func(e *employeeRepository) Delete(id string) error {
	_,err := e.db.Exec(constant.EMPLOYEE_DELETE,id)
	if err != nil {
		return fmt.Errorf("repo : Error Delete Employee : %s ", err.Error())
	}
	return nil
}

// Constructornya
func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{
		db: db,
	}
}