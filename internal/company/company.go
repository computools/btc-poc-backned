package company

import (
	"context"
)

type Service struct {
	db Database
}

func (s *Service) GetCompanies(ctx context.Context) ([]Company, error) {
	companies, err := s.db.GetCompanies(ctx)
	return companiesFromDatabase(companies), err
}

func (s *Service) GetCompany(ctx context.Context, id int64) (Company, error) {
	company, err := s.db.GetCompany(ctx, id)
	return companyFromDatabase(company), err
}

func (s *Service) CreateCompany(ctx context.Context, company Company) (Company, error) {
	c, err := s.db.CreateCompany(ctx, company.toDatabase())
	return companyFromDatabase(c), err
}

func (s *Service) UpdateCompany(ctx context.Context, company Company) (Company, error) {
	c, err := s.db.UpdateCompany(ctx, company.toDatabase())
	return companyFromDatabase(c), err
}

func (s *Service) DeleteCompany(ctx context.Context, id int64) error {
	return s.db.DeleteCompany(ctx, id)
}

func NewService(db Database) *Service {
	return &Service{db: db}
}
