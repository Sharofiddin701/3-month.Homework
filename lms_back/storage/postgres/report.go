package postgres

import (
	"context"
	"lms_back/api/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ReportAdmin struct {
	db *pgxpool.Pool
}

func ReportAdminPayment(db *pgxpool.Pool) ReportAdmin {
	return ReportAdmin{
		db: db,
	}
}
func (c *ReportAdmin) Get(ctx context.Context, report models.AdminPrimaryKey) (models.ReportAdmin, error) {
	reportAdmin := models.ReportAdmin{}

	query := `SELECT 
	a.id AS admin_id,
	a.full_name AS admin_fullname,
	p.id AS payment_id,
	p.price,
	p.student_id,
	p.branch_id,
	p.created_at AS payment_created_at,
	p.updated_at AS payment_updated_at
FROM 
	admin a
JOIN 
	payment p ON a.id = p.admin_id
WHERE 
	a.id = $1`

	row := c.db.QueryRow(context.Background(), query, report.Id)

	err := row.Scan(&reportAdmin.Admin_id,
		&reportAdmin.Full_Name,
		&reportAdmin.PaymentDetail.Payment_id,
		&reportAdmin.PaymentDetail.Price,
		&reportAdmin.PaymentDetail.Student_id,
		&reportAdmin.PaymentDetail.Branch_id,
		&reportAdmin.PaymentDetail.CreatedAt,
		&reportAdmin.PaymentDetail.UpdatedAt)

	if err != nil {
		return models.ReportAdmin{}, err
	}

	return reportAdmin, nil
}
