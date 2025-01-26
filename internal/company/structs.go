package company

import "btc-backend/pkg/database"

type Company struct {
	ID              int64  `json:"id"`
	PublicID        string `json:"public_id"`
	Name            string `json:"name"`
	ShortName       string `json:"short_name"`
	PhysicalAddress string `json:"physical_address"`
	PhysicalCity    string `json:"physical_city"`
	PhysicalZip     string `json:"physical_zip"`
	PostalAddress   string `json:"postal_address"`
	PostalCity      string `json:"postal_city"`
	PostalZip       string `json:"postal_zip"`
}

func (c *Company) toDatabase() database.Company {
	return database.Company{
		ID:              c.ID,
		PublicID:        c.PublicID,
		Name:            c.Name,
		ShortName:       c.ShortName,
		PhysicalAddress: c.PhysicalAddress,
		PhysicalCity:    c.PhysicalCity,
		PhysicalZip:     c.PhysicalZip,
		PostalAddress:   c.PostalAddress,
		PostalCity:      c.PostalCity,
		PostalZip:       c.PostalZip,
	}
}

func companyFromDatabase(c database.Company) Company {
	return Company{
		ID:              c.ID,
		PublicID:        c.PublicID,
		Name:            c.Name,
		ShortName:       c.ShortName,
		PhysicalAddress: c.PhysicalAddress,
		PhysicalCity:    c.PhysicalCity,
		PhysicalZip:     c.PhysicalZip,
		PostalAddress:   c.PostalAddress,
		PostalCity:      c.PostalCity,
		PostalZip:       c.PostalZip,
	}
}

func companiesFromDatabase(c []database.Company) []Company {
	result := make([]Company, 0, len(c))
	for _, c := range c {
		result = append(result, companyFromDatabase(c))
	}

	return result
}
