package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"rent-car/api/models"
	"rent-car/pkg"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrder(db *pgxpool.Pool) orderRepo {
	return orderRepo{
		db: db,
	}
}

func (o *orderRepo) CreateOrder(ctx context.Context, order models.CreateOrderr) (string, error) {
	id := uuid.New()

	query := `insert into orders(
		id,
		car_id,
		customer_id,
		from_date,
		to_date,
		status,
		payment_status,
		amount,
		created_at) values($1,$2,$3,$4,$5,$6,$7,$8,CURRENT_TIMESTAMP)`

	_, err := o.db.Exec(context.Background(), query, id.String(), order.CarId, order.CustomerId, order.FromDate, order.ToDate, order.Status, order.Payment_status, order.Amount)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (o *orderRepo) UpdateOrder(ctx context.Context, order models.GetOrder) (string, error) {
	query := `update orders set 
        customer_id = $1,
        car_id = $2,
        from_date = $3,
        to_date = $4,
        status = $5,
        payment_status = $6,
        amount = $7,
        updated_at = CURRENT_TIMESTAMP
        where id = $8`

	_, err := o.db.Exec(context.Background(), query, order.Customer.Id, order.Car.Id, order.FromDate, order.ToDate, order.Status, order.Status, order.Payment_status, order.Id)
	if err != nil {
		return "", err
	}
	return order.Customer.Id, nil
}

func (o *orderRepo) GetOne(ctx context.Context, id string) (models.GetOrder, error) {
	var resp models.GetOrder

	query := `SELECT
		o.id,
		c.name AS car_name,
		c.id AS car_id,
		c.brand AS car_brand,
		cu.id AS customer_id,
		cu.first_name AS customer_first_name,
		cu.gmail AS customer_email,
		cu.phone AS customer_number,
		o.from_date,
		o.to_date,
		o.status,
		o.payment_status,
		o.amount,
		o.created_at,
		o.updated_at
		FROM orders o
		JOIN cars c ON o.car_id = c.id
		JOIN customers cu ON o.customer_id = cu.id
		WHERE o.id = $1`

	row := o.db.QueryRow(context.Background(), query, id)

	order := models.GetOrder{
		Car:      models.Car{},
		Customer: models.Customer{},
	}

	err := row.Scan(
		&order.Id,
		&order.Car.Name,
		&order.Car.Id,
		&order.Car.Brand,
		&order.Customer.Id,
		&order.Customer.FirstName,
		&order.Customer.Gmail,
		&order.Customer.Phone,
		&order.FromDate,
		&order.ToDate,
		&order.Status,
		&order.Payment_status,
		&order.Amount,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = order

	return resp, nil
}

func (o *orderRepo) GetAll(ctx context.Context, req models.GetAllOrdersRequest) (models.GetAllOrdersResponse, error) {
	var (
		resp   = models.GetAllOrdersResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit
	if req.Search != "" {
		filter += fmt.Sprintf(`and  status ILIKE '%%%v%%'`, req.Search)
	}
	filter += fmt.Sprintf("OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter:", filter)

	query := `Select 
		o.id,
		o.from_date,
		o.to_date,
		o.status,
		o.amount,
		o.created_at,
		o.updated_at,
		c.id as car_id,
		c.name as car_name,
		c.brand as car_brand,
		c.engine_cap as car_engine_cap,
		cu.id as customer_id,
		cu.first_name as customer_first_name,
		cu.last_name as customer_last_name,
		cu.gmail as customer_gmail,
		cu.phone as customer_phone
		From orders o JOIN cars c ON o.car_id = c.id
		JOIN customers cu ON o.customer_id = cu.id 	`
	rows, err := o.db.Query(context.Background(), query+filter+``)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			order = models.GetOrder{
				Car:      models.Car{},
				Customer: models.Customer{},
			}
			updateAt sql.NullString
		)

		err := rows.Scan(
			&order.Id,
			&order.FromDate,
			&order.ToDate,
			&order.Status,
			&order.Amount,
			&order.CreatedAt,
			&updateAt,
			&order.Car.Id,
			&order.Car.Name,
			&order.Car.Brand,
			&order.Car.EngineCap,
			&order.Customer.Id,
			&order.Customer.FirstName,
			&order.Customer.LastName,
			&order.Customer.Gmail, &order.Customer.Phone)
		if err != nil {
			return resp, err
		}
		order.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Orders = append(resp.Orders, order)
	}
	if err = rows.Err(); err != nil {
		return resp, err
	}
	countQuery := `SELECT COUNT(*) FROM orders`

	err = o.db.QueryRow(context.Background(), countQuery).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (o *orderRepo) DeleteOrder(ctx context.Context, id string) error {

	query := ` UPDATE  set
			deleted_at = date_part('epoch', CURRENT_TIMESTAMP)::int
		WHERE id = $1 AND deleted_at=0
	`

	_, err := o.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}
