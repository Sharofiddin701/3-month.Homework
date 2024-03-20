package postgres

import (
	"database/sql"
	"fmt"
	"clone/rent_car_us/api/models"
	"clone/rent_car_us/pkg"
	"github.com/google/uuid"
)

type customerRepo struct {
	db *sql.DB
}

func NewCustomer(db *sql.DB) customerRepo {
	return customerRepo{
		db: db,
	}
}

/*
create (body) id,err
update (body) id,err
delete (id) err
get (id) body,err
getAll (search) []body,count,err
*/

func (c *customerRepo) Create(customer models.Customer) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO customers (
		id ,       
		first_name,
		last_name ,
		gmail ,    
		phone     )
		VALUES($1,$2,$3,$4,$5) `

	_, err := c.db.Exec(query,
		id.String(),
		customer.First_name, customer.Last_name,
		customer.Gmail, customer.Phone)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (c *customerRepo) Update(customer models.Customer) (string, error) {

	query := ` UPDATE customers set
	        first_name=$1,
	        last_name=$2,
	        gmail=$3,
	        phone=$4,
			updated_at=CURRENT_TIMESTAMP
		WHERE id = $5 
	`

	_, err := c.db.Exec(query,
		customer.First_name, customer.Last_name,
		customer.Gmail, customer.Phone,customer.Id)

	if err != nil {
		return "", err
	}

	return customer.Id, nil
}

func (c *customerRepo) GetAllCustomers(req models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error) {
	var (
		resp   = models.GetAllCustomersResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit
	if req.Search != "" {
		filter += fmt.Sprintf(` and first_name ILIKE '%%%v%%'`, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter:", filter)

	query:=`SELECT
    cu.id AS customer_id,
    cu.first_name AS customer_first_name,
    cu.last_name AS customer_last_name,
    cu.gmail AS customer_gmail,
    cu.phone AS customer_phone,
    COUNT(o.id) AS order_count,
    COUNT(ca.id) AS car_count,
    cu.created_at,
    cu.updated_at,
    o.id AS order_id,
    o.from_date,
    o.to_date,
    o.status,
    o.paid,
    o.amount
FROM
    customers cu
JOIN
    orders o ON cu.id = o.customer_id
JOIN
    cars ca ON ca.id = o.car_id
GROUP BY
    cu.id,
    cu.first_name,
    cu.last_name,
    cu.gmail,
    cu.phone,
    cu.created_at,
    cu.updated_at,
    o.id,
    o.from_date,
    o.to_date,
    o.status,
    o.paid,
    o.amount;`

	rows, err := c.db.Query(query )
	if err != nil {
		return resp, err
	}
     defer rows.Close()

	for rows.Next() {
		var (
			customer = models.GetAllCustomer{
				Order: models.GetOrder{},
			}
			updateAt sql.NullString
		)
		if err := rows.Scan(
			&customer.Id,
			&customer.FirstName,
			&customer.LastName,
			&customer.Gmail,
			&customer.Phone,
			&customer.OrderCount,
			&customer.CarsCount,
			&customer.CreatedAt,
			&updateAt,
			&customer.Order.Id,       
			&customer.Order.FromDate,       
			&customer.Order.ToDate,       
			&customer.Order.Status,       
			&customer.Order.Paid,       
			&customer.Order.Amount);err != nil {
			return resp, err
		}
		customer.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Customer = append(resp.Customer,models.GetAllCustomer{})
	}
	if err = rows.Err();err != nil {
		return resp,err
	}
	countQuery := `Select count(*) from customers`
	err = c.db.QueryRow(countQuery).Scan(&resp.Count)
	if err != nil {
		return resp,err
	}
	return resp, nil
}

func (c *customerRepo) GetByID(id string) (models.Customer, error) {
	customer := models.Customer{}

	if err := c.db.QueryRow(`select id,
	first_name,
	last_name,gmail,
	phone
	from customers where id = $1`, id).Scan(
		&customer.Id,
		&customer.First_name,
		&customer.Last_name,
		&customer.Gmail,
		&customer.Phone); err != nil {
		return models.Customer{}, err
	}
	return customer, nil
}
func (c *customerRepo) Delete(id string) error {

	query := ` UPDATE customers set
			deleted_at = date_part('epoch', CURRENT_TIMESTAMP)::int
		WHERE id = $1 
	`

	_, err := c.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}



func (c *customerRepo) GetAllCustomerCars(req models.GetAllCustomerCarsRequest) (models.GetAllCustomerCarsResponse, error) {
	var (
		resp   = models.GetAllCustomerCarsResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit
	if req.Search != "" {
		filter += fmt.Sprintf(` and ca.name ILIKE '%%%v%%'`, req.Search)
	}

	if req.Id != "" {
		filter += fmt.Sprintf(` Where cu.id= '%%%v%%'`, req.Id)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter:", filter)

	query:=`Select 
	cu.id as customer_id,
	ca.name as car_name,
	o.created_ad as creatad_at,
	o.amount as price
	From customers cu JOIN orders o ON  cu.id = o.customer_id Join cars ca  ON ca.id=o.car_id 
	`

	rows, err := c.db.Query(query + filter + ``)
	if err != nil {
		return resp, err
	}
     defer rows.Close()

	for rows.Next() {
		var (
			customer = models.GetAllCustomerCars{
			}
		)
		if err := rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.CreatedAt,
			&customer.Amount);err != nil {
			return resp, err
		}
		resp.Customer = append(resp.Customer,customer)
	}
	if err = rows.Err();err != nil {
		return resp,err
	}
	countQuery := `Select count(*) from customers`
	err = c.db.QueryRow(countQuery).Scan(&resp.Count)
	if err != nil {
		return resp,err
	}
	return resp, nil
}
