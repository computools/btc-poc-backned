package financialreports

import (
	"context"
)

type Service struct {
	db Database
}

func (s *Service) GetFinancialReport(ctx context.Context, id int64) (FinancialReports, error) {
	financialReport, err := s.db.GetFinancialReport(ctx, id)
	return financialReportFromDatabase(financialReport), err
}

func (s *Service) GetFinancialReportsByCompanyID(ctx context.Context, id int64) ([]FinancialReports, error) {
	financialReports, err := s.db.GetFinancialReportsByCompanyID(ctx, id)
	return financialReportsFromDatabase(financialReports), err
}

func (s *Service) CreateFinancialReport(ctx context.Context, financialReport FinancialReports) (FinancialReports, error) {
	fr, err := s.db.CreateFinancialReport(ctx, financialReport.toDatabase())
	return financialReportFromDatabase(fr), err
}

func (s *Service) UpdateFinancialReport(ctx context.Context, financialReport FinancialReports) (FinancialReports, error) {
	fr, err := s.db.UpdateFinancialReport(ctx, financialReport.toDatabase())
	return financialReportFromDatabase(fr), err
}

func (s *Service) DeleteFinancialReport(ctx context.Context, id int64) error {
	return s.db.DeleteFinancialReport(ctx, id)
}

func NewService(db Database) *Service {
	return &Service{db: db}
}
