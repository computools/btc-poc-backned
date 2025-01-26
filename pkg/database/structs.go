package database

type Company struct {
	ID              int64
	PublicID        string
	Name            string
	ShortName       string
	PhysicalAddress string
	PhysicalCity    string
	PhysicalZip     string
	PostalAddress   string
	PostalCity      string
	PostalZip       string
}

func (Company) TableName() string {
	return "company"
}

type User struct {
	ID         int64
	FirstName  string
	LastName   string
	Position   string
	CompanyID  *int64
	ImageURL   string
	KeycloakID string
}

func (User) TableName() string {
	return "user"
}

type FinancialReports struct {
	ID        int64
	Name      string
	Year      int64
	CompanyID int64
	URL       string
}

func (FinancialReports) TableName() string {
	return "financial_reports"
}

type CompanyPerformance struct {
	ID                     int64
	CompanyID              int64
	Quarter                int64
	Year                   int64
	Income                 float64
	TotalWageBill          float64
	ExecutiveWageBill      float64
	TotalStaff             int64
	TotalExecutive         int64
	TotalAssets            float64
	OperationalExpenditure float64
}

func (CompanyPerformance) TableName() string {
	return "company_performance"
}
