package database

import (
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/resq/internal/models"
)

func CreateCompany(userId string, company *models.CreateCompanyBody) (createdCompanyId string, err error) {
	nanoid, err := gonanoid.New()
	if err != nil {
		return "", err
	}
	userFromDb, err := GetUserById(userId)
	if err != nil {
		return "", err
	}
	if userFromDb.Role == "ngo_admin" || userFromDb.Role == "company_admin" {
		return "", errors.New("user already has a ngo / company")
	}
	err = db.QueryRow("INSERT INTO companies (id, name, description, creator_id) VALUES ($1, $2, $3, $4) RETURNING id", nanoid, company.Name, company.Description, userId).
		Scan(&createdCompanyId)
	if err != nil {
		return "", err
	}
	_, err = db.Exec("UPDATE users SET org_id = $1, role = 'company_admin' WHERE id = $2", createdCompanyId, userId)
	if err != nil {
		return "", err
	}
	_, err = db.Exec("INSERT INTO companies_admins (user_id, company_id) VALUES ($1, $2)", userId, createdCompanyId)
	if err != nil {
		return "", err
	}
	return createdCompanyId, nil
}

func GetCompaniesAdmins() (companiesAdmins []models.CompanyAdmin, err error) {
	rows, err := db.Query("SELECT * FROM companies_admins")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var companyAdmin models.CompanyAdmin
		if err := rows.Scan(&companyAdmin.UserId, &companyAdmin.CompanyId); err != nil {
			return nil, err
		}
		companiesAdmins = append(companiesAdmins, companyAdmin)
	}
	return companiesAdmins, nil
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

func GetUserCompany(userId string) (company models.Company, err error) {
	if err := db.QueryRow(`
	SELECT c.id AS company_id, c.name AS company_name, c.description AS company_description, c.rating AS company_rating
	FROM companies c
	WHERE c.creator_id = $1`, userId).
		Scan(&company.Id, &company.Name, &company.Description, &company.Rating); err != nil {
		return company, err
	}
	return company, err
}
