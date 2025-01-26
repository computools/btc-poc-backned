package company

import (
	"context"

	"btc-backend/pkg/database"
)

type Database interface {
	GetCompany(ctx context.Context, id int64) (database.Company, error)
	GetCompanies(ctx context.Context) ([]database.Company, error)
	CreateCompany(ctx context.Context, company database.Company) (database.Company, error)
	UpdateCompany(ctx context.Context, company database.Company) (database.Company, error)
	DeleteCompany(ctx context.Context, id int64) error
}
