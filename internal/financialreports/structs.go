package financialreports

import "btc-backend/pkg/database"

type FinancialReports struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Year       int64  `json:"year"`
	CompanyID  int64  `json:"company_id"`
	URL        string `json:"url"`
	PreviewURL string `json:"preview_url"`
}

func (f *FinancialReports) toDatabase() database.FinancialReports {
	return database.FinancialReports{
		ID:         f.ID,
		Name:       f.Name,
		Year:       f.Year,
		CompanyID:  f.CompanyID,
		URL:        f.URL,
		PreviewURL: f.PreviewURL,
	}
}

func financialReportFromDatabase(f database.FinancialReports) FinancialReports {
	return FinancialReports{
		ID:         f.ID,
		Name:       f.Name,
		Year:       f.Year,
		CompanyID:  f.CompanyID,
		URL:        f.URL,
		PreviewURL: f.PreviewURL,
	}
}

func financialReportsFromDatabase(f []database.FinancialReports) []FinancialReports {
	result := make([]FinancialReports, 0, len(f))
	for _, f := range f {
		result = append(result, financialReportFromDatabase(f))
	}

	return result
}
