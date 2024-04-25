package models

type ReportAdmin struct {
	Admin_id      string `json:"admin_id"`
	Full_Name     string `json:"full_name"`
	PaymentDetail ReportAdminPayment
}

type ReportAdminPayment struct {
	Payment_id string  `json:"payment_id"`
	Price      float64 `json:"price"`
	Student_id string  `json:"student_id"`
	Branch_id  string  `json:"branch_id"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}
