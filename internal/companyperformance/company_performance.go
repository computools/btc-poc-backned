package companyperformance

import (
	"context"
)

type Service struct {
	db Database
}

func (s *Service) GetCompanyPerformance(ctx context.Context, id int64) (CompanyPerformance, error) {
	companyPerformance, err := s.db.GetCompanyPerformance(ctx, id)
	return companyPerformanceFromDatabase(companyPerformance), err
}

func (s *Service) GetCompanyPerformancesByCompanyID(ctx context.Context, id int64) ([]CompanyPerformance, error) {
	companyPerformances, err := s.db.GetCompanyPerformancesByCompanyID(ctx, id)
	return companyPerformancesFromDatabase(companyPerformances), err
}

func (s *Service) CreateCompanyPerformance(ctx context.Context, companyPerformance CompanyPerformance) (CompanyPerformance, error) {
	cp, err := s.db.CreateCompanyPerformance(ctx, companyPerformance.toDatabase())
	return companyPerformanceFromDatabase(cp), err
}

func (s *Service) UpdateCompanyPerformance(ctx context.Context, companyPerformance CompanyPerformance) (CompanyPerformance, error) {
	cp, err := s.db.UpdateCompanyPerformance(ctx, companyPerformance.toDatabase())
	return companyPerformanceFromDatabase(cp), err
}

func (s *Service) DeleteCompanyPerformance(ctx context.Context, id int64) error {
	return s.db.DeleteCompanyPerformance(ctx, id)
}

func NewService(db Database) *Service {
	return &Service{db: db}
}
