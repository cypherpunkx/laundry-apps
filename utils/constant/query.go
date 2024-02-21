package constant

const (
	/// ====================== Master Data Employee ======================
	EMPLOYEE_INSERT =  "INSERT INTO employee(id,name,phone_number,address)VALUES($1, $2, $3, $4)"
	EMPLOYEE_LIST = "SELECT * FROM employee"
	EMPLOYEE_GET = "SELECT * FROM employee where id=$1"
	EMPLOYEE_UPDATE = "UPDATE employee SET name=$1, phone_number=$2, address=$3 WHERE id=$4"
	EMPLOYEE_DELETE = "DELETE FROM employee WHERE id=$1"

	/// ====================== Master Data Product ======================
	PRODUCT_INSERT = "INSERT INTO product (id,name,price,uom) VALUES ($1,$2,$3,$4)"
	PRODUCT_GET = "SELECT * FROM product where id=$1"
	PRODUCT_UPDATE = "UPDATE product SET name=$1,price=$2,uom=$3 WHERE id=$4"
	PRODUCT_DELETE = "DELETE FROM product WHERE id=$1"
	/// ====================== Master Data Customer ======================
	CUSTOMER_INSERT =  "INSERT INTO customer(id,name,phone_number,address)VALUES($1, $2, $3, $4)"
	CUSTOMER_GET = "SELECT * FROM customer where id=$1"
	//...
	/// ====================== Data Bill ======================
	BILL_CREATE = "INSERT INTO bill (id,bill_date,entry_date,employee_id,customer_id) values ($1,$2,$3,$4,$5)"
	BILL_GET = "SELECT * FROM bill where id=$1"
	//...
	/// ====================== Data Bill Details ======================
	BIll_DETAIL_CREATE = "INSERT INTO bill_detail (id,bill_id,product_id,product_price,qty,finish_date) values ($1,$2,$3,$4,$5,$6)"
	BIll_DETAIL_GET = "SELECT * FROM bill_detail WHERE bill_id=$1"
	/// ====================== Data User Picture ======================
	CREATE_USER_PICTURE = "INSERT INTO user_picture(id,user_id,file_location) values ($1,$2,$3)"
	GET_USER_PICTURE = "SELECT * FROM user_picture where user_id=$1"
)

