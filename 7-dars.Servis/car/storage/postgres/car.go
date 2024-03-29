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

type carRepo struct {
	db *pgxpool.Pool
}

func NewCar(db *pgxpool.Pool) carRepo {
	return carRepo{
		db: db,
	}
}

func (c *carRepo) GetAvailableCars(ctx context.Context, req models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {
	var (
		resp   = models.GetAllCarsResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter:", filter)
	rows, err := c.db.Query(context.Background(), `
	SELECT id, name, model, year, brand,
	hourse_power, colour, engine_cap, created_at
	FROM cars
	WHERE id NOT IN (SELECT car_id FROM orders WHERE car_id IS NOT NULL)`+filter+``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			car = models.Car{}
			created_at sql.NullString
		)

		if err := rows.Scan(
			&car.Id,
			&car.Name,
			&car.Model,
			&car.Year,
			&car.Brand,
			&car.HoursePower,
			&car.Colour,
			&car.EngineCap,
			&created_at);
		    err != nil {
			return resp, err
		}
		resp.Cars = append(resp.Cars, car)
	}
	return resp, nil
}

func (c *carRepo) Create(ctx context.Context, car models.CreateCar) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO cars (
		id,
		name,
		brand,
		model,
		hourse_power,
		colour,
		engine_cap,
		year)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8) 
	`

	_, err := c.db.Exec(context.Background(), query,
		id.String(),
		car.Name, car.Brand,
		car.Model, car.HoursePower,
		car.Colour, car.EngineCap, car.Year)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (c *carRepo) Update(ctx context.Context, car models.Car) (string, error) {

	query := `UPDATE cars set
			name=$1,
			brand=$2,
			model=$3,
			hourse_power=$4,
			colour=$5,
			engine_cap=$6,
			updated_at=CURRENT_TIMESTAMP
		WHERE id = $7 AND deleted_at = 0
	`

	_, err := c.db.Exec(context.Background(), query,
		car.Name, car.Brand,
		car.Model, car.HoursePower,
		car.Colour, car.EngineCap, car.Id)

	if err != nil {
		return "", err
	}

	return car.Id, nil
}

func (c *carRepo) GetAll(ctx context.Context, req models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {
	var (
		resp   = models.GetAllCarsResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf("OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter:", filter)
	rows, err := c.db.Query(context.Background(), `select 
				count(id) OVER(),
				id, 
				name,
				brand,
				model,
				year,
				hourse_power,
				colour,
				engine_cap,
				--created_at::date,
				updated_at,
				year
	  FROM cars WHERE deleted_at = 0 `+filter+``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			car      = models.Car{}
			updateAt sql.NullString
		)

		if err := rows.Scan(
			&resp.Count,
			&car.Id,
			&car.Name,
			&car.Brand,
			&car.Model,
			&car.Year,
			&car.HoursePower,
			&car.Colour,
			&car.EngineCap,
			// &car.CreatedAt,
			&updateAt,
			&car.Year); err != nil {
			return resp, err
		}

		car.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Cars = append(resp.Cars, car)
	}
	return resp, nil
}

func (c *carRepo) GetByID(ctx context.Context, id string) (models.Car, error) {
	car := models.Car{}
	var CreatedAt sql.NullString
	if err := c.db.QueryRow(context.Background(), `select id, name, brand, model, hourse_power, colour, engine_cap, created_at, year from cars where id = $1`, id).Scan(
		&car.Id,
		&car.Name,
		&car.Brand,
		&car.Model,
		&car.HoursePower,
		&car.Colour,
		&car.EngineCap,
		&CreatedAt,
		&car.Year,
	); err != nil {
		return models.Car{}, err
	}
	return car, nil
}

func (c *carRepo) Delete(ctx context.Context, id string) error {

	query := `delete from cars WHERE id = $1`
	_, err := c.db.Exec(context.Background(), query, id)

	if err != nil {
		return err
	}

	return nil
}

func (c *carRepo) GetAllOrdersCars(req models.GetAllCarsRequest) ([]models.GetOrder, error) {
	var orders []models.GetOrder

	offset := (req.Page - 1) * req.Limit

	query := `SELECT DISTINCT
                o.id,
				c.id AS car_id,
                c.name AS car_name,
                c.brand AS car_brand,
				c.model as car_model,
                cu.id AS customer_id,
                cu.first_name AS customer_first_name,
                o.from_date,
                o.to_date,
                o.amount,
                o.created_at,
                o.updated_at 
            FROM 
                orders o
            JOIN 
                cars c ON o.car_id = c.id
            JOIN 
                customers cu ON o.customer_id = cu.id
            WHERE 1=1`

	if req.Search != "" {
		query += fmt.Sprintf(`AND c.name ILIKE '%%%s%%'`, req.Search)
	}
	query += fmt.Sprintf("ORDER BY o.created_at OFFSET %d LIMIT %d", offset, req.Limit)

	rows, err := c.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.GetOrder
		err := rows.Scan(
			&order.Id,
			&order.Car.Id,
			&order.Car.Name,
			&order.Car.Brand,
			&order.Car.Model,
			&order.Customer.Id,
			&order.Customer.FirstName,
			&order.FromDate,
			&order.ToDate,
			&order.Amount,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (c *carRepo) GetByIDCustomerCarr(id string) (models.GetOrder, error) {
	order := models.GetOrder{
		Car:      models.Car{},
		Customer: models.Customer{},
	}

	if err := c.db.QueryRow(context.Background(), `SELECT 
            o.id,
            c.id AS car_id,
            c.name AS car_name,
            c.brand AS car_brand,
            c.model as car_model,
            cu.id AS customer_id,
            cu.first_name AS customer_first_name,
            cu.gmail AS customer_gmail,
            o.from_date,
            o.to_date,
            o.status,
            o.amount,
            o.created_at,
            o.updated_at 
            FROM 
                orders o
            JOIN 
                cars c ON o.car_id = c.id
            JOIN 
                customers cu ON o.customer_id = cu.id
            WHERE c.id=$1`, id).Scan(
		&order.Id,
		&order.Car.Id,
		&order.Car.Name,
		&order.Car.Brand,
		&order.Car.Model,
		&order.Customer.Id,
		&order.Customer.FirstName,
		&order.Customer.Gmail,
		&order.FromDate,
		&order.ToDate,
		&order.Status,
		&order.Amount,
		&order.CreatedAt,
		&order.UpdatedAt,
	); err != nil {
		return models.GetOrder{}, err
	}
	return order, nil
}
