package database

import (
	"context"
	"errors"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	glog "gorm.io/gorm/logger"

	"btc-backend/config"
)

type Postgres struct {
	cfg *config.DatabaseConfig
	db  *gorm.DB
}

func (p *Postgres) Ping(ctx context.Context) error {
	db, err := p.db.DB()
	if err != nil {
		return err
	}

	return db.PingContext(ctx)
}

func (p *Postgres) GetCompany(ctx context.Context, id int64) (Company, error) {
	var company Company
	err := p.db.WithContext(ctx).Model(&Company{}).Where("id = ?", id).First(&company).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Company{}, ErrRecordNotFound
	}

	return company, nil
}

func (p *Postgres) GetCompanies(ctx context.Context) ([]Company, error) {
	var company []Company
	err := p.db.WithContext(ctx).Model(&Company{}).Find(&company).Error
	return company, err
}

func (p *Postgres) CreateCompany(ctx context.Context, company Company) (Company, error) {
	err := p.db.WithContext(ctx).Model(&company).Clauses(clause.Returning{}).Create(&company).Error
	return company, err
}

func (p *Postgres) UpdateCompany(ctx context.Context, company Company) (Company, error) {
	err := p.db.WithContext(ctx).Model(&company).Clauses(clause.Returning{}).Where("id = ?", company.ID).Updates(map[string]any{
		"public_id":        company.PublicID,
		"name":             company.Name,
		"short_name":       company.ShortName,
		"physical_address": company.PhysicalAddress,
		"physical_city":    company.PhysicalCity,
		"physical_zip":     company.PhysicalZip,
		"postal_address":   company.PostalAddress,
		"postal_city":      company.PostalCity,
		"postal_zip":       company.PostalZip,
	}).Error
	return company, err
}

func (p *Postgres) DeleteCompany(ctx context.Context, id int64) error {
	return p.db.WithContext(ctx).Model(&Company{}).Where("id = ?", id).Delete(&Company{}).Error
}

func (p *Postgres) GetUser(ctx context.Context, id int64) (User, error) {
	var user User
	err := p.db.WithContext(ctx).Model(&User{}).Where("id = ?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return User{}, ErrRecordNotFound
	}

	return user, nil
}

func (p *Postgres) GetUsersByCompanyID(ctx context.Context, id int64) ([]User, error) {
	var users []User
	err := p.db.WithContext(ctx).Model(&User{}).Where("company_id = ?", id).Find(&users).Error
	return users, err
}

func (p *Postgres) CreateUser(ctx context.Context, user User) (User, error) {
	err := p.db.WithContext(ctx).Model(&user).Clauses(clause.Returning{}).Create(&user).Error
	return user, err
}

func (p *Postgres) UpdateUser(ctx context.Context, user User) (User, error) {
	err := p.db.WithContext(ctx).
		Model(&user).
		Clauses(clause.Returning{}).
		Where("id = ?", user.ID).
		Updates(map[string]any{
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"position":   user.Position,
			"company_id": user.CompanyID,
			"image_url":  user.ImageURL,
		}).Error
	return user, err
}

func (p *Postgres) DeleteUser(ctx context.Context, id int64) error {
	return p.db.WithContext(ctx).Model(&User{}).Where("id = ?", id).Delete(&User{}).Error
}

func (p *Postgres) GetFinancialReport(ctx context.Context, id int64) (FinancialReports, error) {
	var financialReport FinancialReports
	err := p.db.WithContext(ctx).Model(&FinancialReports{}).Where("id = ?", id).First(&financialReport).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return FinancialReports{}, ErrRecordNotFound
	}

	return financialReport, nil
}

func (p *Postgres) GetFinancialReportsByCompanyID(ctx context.Context, id int64) ([]FinancialReports, error) {
	var financialReports []FinancialReports
	err := p.db.WithContext(ctx).Model(&FinancialReports{}).Where("company_id = ?", id).Find(&financialReports).Error
	return financialReports, err
}

func (p *Postgres) CreateFinancialReport(ctx context.Context, financialReport FinancialReports) (FinancialReports, error) {
	err := p.db.WithContext(ctx).Model(&financialReport).Clauses(clause.Returning{}).Create(&financialReport).Error
	return financialReport, err
}

func (p *Postgres) UpdateFinancialReport(ctx context.Context, financialReport FinancialReports) (FinancialReports, error) {
	err := p.db.WithContext(ctx).
		Model(&financialReport).
		Clauses(clause.Returning{}).
		Where("id = ?", financialReport.ID).
		Updates(map[string]any{
			"name": financialReport.Name,
			"year": financialReport.Year,
			"url":  financialReport.URL,
		}).Error
	return financialReport, err
}

func (p *Postgres) DeleteFinancialReport(ctx context.Context, id int64) error {
	return p.db.WithContext(ctx).Model(&FinancialReports{}).Where("id = ?", id).Delete(&FinancialReports{}).Error
}

func (p *Postgres) GetCompanyPerformance(ctx context.Context, id int64) (CompanyPerformance, error) {
	var companyPerformance CompanyPerformance
	err := p.db.WithContext(ctx).Model(&CompanyPerformance{}).Where("id = ?", id).First(&companyPerformance).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return CompanyPerformance{}, ErrRecordNotFound
	}

	return companyPerformance, nil
}

func (p *Postgres) GetCompanyPerformancesByCompanyID(ctx context.Context, id int64) ([]CompanyPerformance, error) {
	var companyPerformances []CompanyPerformance
	err := p.db.WithContext(ctx).Model(&CompanyPerformance{}).Where("company_id = ?", id).Find(&companyPerformances).Error
	return companyPerformances, err
}

func (p *Postgres) CreateCompanyPerformance(ctx context.Context, companyPerformance CompanyPerformance) (CompanyPerformance, error) {
	err := p.db.WithContext(ctx).Model(&companyPerformance).Clauses(clause.Returning{}).Create(&companyPerformance).Error
	return companyPerformance, err
}

func (p *Postgres) UpdateCompanyPerformance(ctx context.Context, companyPerformance CompanyPerformance) (CompanyPerformance, error) {
	err := p.db.WithContext(ctx).
		Model(&companyPerformance).
		Clauses(clause.Returning{}).
		Where("id = ?", companyPerformance.ID).
		Updates(map[string]any{
			"quarter":                 companyPerformance.Quarter,
			"year":                    companyPerformance.Year,
			"income":                  companyPerformance.Income,
			"total_wage_bill":         companyPerformance.TotalWageBill,
			"executive_wage_bill":     companyPerformance.ExecutiveWageBill,
			"total_staff":             companyPerformance.TotalStaff,
			"total_executive":         companyPerformance.TotalExecutive,
			"total_assets":            companyPerformance.TotalAssets,
			"operational_expenditure": companyPerformance.OperationalExpenditure,
		}).Error
	return companyPerformance, err
}

func (p *Postgres) DeleteCompanyPerformance(ctx context.Context, id int64) error {
	return p.db.WithContext(ctx).Model(&CompanyPerformance{}).Where("id = ?", id).Delete(&CompanyPerformance{}).Error
}

func (p *Postgres) Start(_ context.Context) (err error) {
	p.db, err = gorm.Open(postgres.Open(p.cfg.DSN), &gorm.Config{
		Logger: glog.Default.LogMode(glog.Silent),
	})
	if err != nil {
		return
	}

	db, err := p.db.DB()
	if err != nil {
		return
	}

	if err = goose.SetDialect("postgres"); err != nil {
		return
	}

	err = goose.Up(db, p.cfg.MigrationPath, goose.WithAllowMissing())
	return
}

func (p *Postgres) Stop(ctx context.Context) error {
	sql, err := p.db.WithContext(ctx).DB()
	if err != nil {
		return err
	}

	return sql.Close()
}

func NewPostgres(cfg *config.Config) *Postgres {
	return &Postgres{cfg: &cfg.DatabaseConfig}
}
