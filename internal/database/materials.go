package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/resq/internal/models"
)

func CreateMaterial(material *models.CreateMaterial, companyId string) (createdMaterialId string, err error) {
	id, err := gonanoid.New()
	if err != nil {
		return "", err
	}
	if err := db.QueryRow("INSERT INTO materials (id, title, description, company_id) VALUES ($1, $2, $3, $4) RETURNING id",
		id, material.Title, material.Description, companyId).Scan(&createdMaterialId); err != nil {
		return "", err
	}
	return createdMaterialId, nil
}

func GetCompanyMaterials(companyId string) (materials []models.Material, err error) {
	rows, err := db.Query("SELECT * FROM materials WHERE company_id = $1", companyId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var material models.Material
		if err := rows.Scan(&material.Id, &material.Title, &material.Description, &material.CreatedAt, &material.IsActive, &material.CompanyId); err != nil {
			return nil, err
		}
		materials = append(materials, material)
		if err := rows.Err(); err != nil {
			return nil, err
		}
	}

	return materials, nil
}
