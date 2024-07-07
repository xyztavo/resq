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
	err = db.QueryRow("INSERT INTO companies (id, name, description, creator_id) VALUES ($1, $2, $3, $4) RETURNING id", nanoid, company.Name, company.Description, userId).Scan(&createdCompanyId)
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
		rows.Scan(&company.Id, &company.Name, &company.Description, &company.Rating, &company.CreatorId)
		companies = append(companies, company)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	rows.Close()
	return companies, nil
}
