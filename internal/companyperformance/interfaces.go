package companyperformance

import (
	"context"

	"btc-backend/pkg/database"
)

type Database interface {
	GetCompanyPerformance(ctx context.Context, id int64) (database.CompanyPerformance, error)
	GetCompanyPerformancesByCompanyID(ctx context.Context, id int64) ([]database.CompanyPerformance, error)
	CreateCompanyPerformance(ctx context.Context, companyPerformance database.CompanyPerformance) (database.CompanyPerformance, error)
	UpdateCompanyPerformance(ctx context.Context, companyPerformance database.CompanyPerformance) (database.CompanyPerformance, error)
	DeleteCompanyPerformance(ctx context.Context, id int64) error
}
