package postgres

import (
	"context"
	"fmt"
	"rent-car/api/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type cusRepo struct {
	db *pgxpool.Pool
}

func NewCus(db *pgxpool.Pool) cusRepo {
	return cusRepo{
		db: db,
	}
}

func (c *cusRepo) GetCusCars(id string, req models.GetAllCarsRequest) ([]models.GetOrder, error) {
	var orders []models.GetOrder

	offset := (req.Page - 1) * req.Limit

	query := `SELECT DISTINCT
			o.id,
			c.name AS car_name,
			cu.first_name AS customer_first_name,
			o.from_date,
			o.to_date,
			o.created_at
			FROM 
				orders o
			JOIN 
				cars c ON o.car_id = c.id
			JOIN 
				customers cu ON o.customer_id = cu.id
			WHERE cu.id = $1`

	if req.Search != "" {
		query += fmt.Sprintf(` AND c.name ILIKE '%%%s%%'`, req.Search)
	}
	query += fmt.Sprintf(" ORDER BY o.created_at OFFSET %d LIMIT %d", offset, req.Limit)

	rows, err := c.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.GetOrder
		err := rows.Scan(
			&order.Id,
			&order.Car.Name,
			&order.Customer.FirstName,
			&order.FromDate,
			&order.ToDate,
			&order.CreatedAt,
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
