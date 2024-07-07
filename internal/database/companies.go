package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/resq/internal/models"
)

func CreateCompany(userId string, company *models.CreateCompanyBody) (createdCompanyId string, err error) {
	nanoid, err := gonanoid.New()
	if err != nil {
		return "", err
	}
	err = db.QueryRow("INSERT INTO companies (id, name, description, creator_id) VALUES ($1, $2, $3, $4) RETURNING id", nanoid, company.Name, company.Description, userId).
		Scan(&createdCompanyId)
	if err != nil {
		return "", err
	}
	_, err = db.Exec("UPDATE users SET org_type = $1, org_id = $2 WHERE id = $3", "company", createdCompanyId, userId)
	if err != nil {
		return "", err
	}
	return createdCompanyId, nil
}

func GetCompanies() (companies []models.Company, err error) {
	rows, err := db.Query("SELECT * FROM companies")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var company models.Company
		if err := rows.Scan(&company.Id, &company.Name, &company.Description, &company.Rating, &company.CreatorId); err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}

	rows.Close()
	return companies, nil
}

func GetUserAdminCompany(userId string) (company models.Company, err error) {
	if err := db.QueryRow("SELECT * FROM companies WHERE creator_id = $1", userId).
		Scan(&company.Id, &company.Name, &company.Description, &company.Rating, &company.CreatorId); err != nil {
		return company, err
	}
	return company, nil
}

func GetUserCompany(companyId *string) (company models.Company, err error) {
	if err := db.QueryRow("SELECT * FROM companies WHERE id = $1", companyId).
		Scan(&company.Id, &company.Name, &company.Description, &company.Rating, &company.CreatorId); err != nil {
		return company, err
	}
	return company, err
}
