package companyperformance

import "btc-backend/pkg/database"

type CompanyPerformance struct {
	ID                     int64   `json:"id"`
	CompanyID              int64   `json:"company_id"`
	Quarter                int64   `json:"quarter"`
	Year                   int64   `json:"year"`
	Income                 float64 `json:"income"`
	TotalWageBill          float64 `json:"total_wage_bill"`
	ExecutiveWageBill      float64 `json:"executive_wage_bill"`
	TotalStaff             int64   `json:"total_staff"`
	TotalExecutive         int64   `json:"total_executive"`
	TotalAssets            float64 `json:"total_assets"`
	OperationalExpenditure float64 `json:"operational_expenditure"`
}

func (cp *CompanyPerformance) toDatabase() database.CompanyPerformance {
	return database.CompanyPerformance{
		ID:                     cp.ID,
		CompanyID:              cp.CompanyID,
		Quarter:                cp.Quarter,
		Year:                   cp.Year,
		Income:                 cp.Income,
		TotalWageBill:          cp.TotalWageBill,
		ExecutiveWageBill:      cp.ExecutiveWageBill,
		TotalStaff:             cp.TotalStaff,
		TotalExecutive:         cp.TotalExecutive,
		TotalAssets:            cp.TotalAssets,
		OperationalExpenditure: cp.OperationalExpenditure,
	}
}

func companyPerformanceFromDatabase(cp database.CompanyPerformance) CompanyPerformance {
	return CompanyPerformance{
		ID:                     cp.ID,
		CompanyID:              cp.CompanyID,
		Quarter:                cp.Quarter,
		Year:                   cp.Year,
		Income:                 cp.Income,
		TotalWageBill:          cp.TotalWageBill,
		ExecutiveWageBill:      cp.ExecutiveWageBill,
		TotalStaff:             cp.TotalStaff,
		TotalExecutive:         cp.TotalExecutive,
		TotalAssets:            cp.TotalAssets,
		OperationalExpenditure: cp.OperationalExpenditure,
	}
}

func companyPerformancesFromDatabase(cp []database.CompanyPerformance) []CompanyPerformance {
	result := make([]CompanyPerformance, 0, len(cp))
	for _, cp := range cp {
		result = append(result, companyPerformanceFromDatabase(cp))
	}

	return result
}
