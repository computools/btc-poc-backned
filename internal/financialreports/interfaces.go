package financialreports

import (
	"context"

	"btc-backend/pkg/database"
)

type Database interface {
	GetFinancialReport(ctx context.Context, id int64) (database.FinancialReports, error)
	GetFinancialReportsByCompanyID(ctx context.Context, id int64) ([]database.FinancialReports, error)
	CreateFinancialReport(ctx context.Context, financialReport database.FinancialReports) (database.FinancialReports, error)
	UpdateFinancialReport(ctx context.Context, financialReport database.FinancialReports) (database.FinancialReports, error)
	DeleteFinancialReport(ctx context.Context, id int64) error
}
