package service

import (
	"context"
	"fmt"
	"lms_back/api/models"
	"lms_back/storage"
)

type ReportAdminService struct {
	storage storage.IStorage
}

func ReportAdminPayment(storage storage.IStorage) ReportAdminService {
	return ReportAdminService{
		storage: storage,
	}
}
func (u ReportAdminService) GetReportAdmin(ctx context.Context, report models.AdminPrimaryKey) (models.ReportAdmin, error) {
	pKey, err := u.storage.Report().Get(ctx, report)
	if err != nil {
		fmt.Println("ERROR in service layer while getting report:", err.Error())
		return models.ReportAdmin{}, err
	}

	return pKey, nil
}
