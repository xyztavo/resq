package database

import (
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/resq/internal/models"
)

func CreateNGO(userId string, company *models.CreateNGOBody) (createdNGOId string, err error) {
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
	err = db.QueryRow("INSERT INTO ngos (id, name, description, creator_id) VALUES ($1, $2, $3, $4) RETURNING id", nanoid, company.Name, company.Description, userId).
		Scan(&createdNGOId)
	if err != nil {
		return "", err
	}
	_, err = db.Exec("UPDATE users SET org_id = $1, role = 'ngo_admin' WHERE id = $2", createdNGOId, userId)
	if err != nil {
		return "", err
	}
	_, err = db.Exec("INSERT INTO ngos_admins (user_id, ngo_id) VALUES ($1, $2)", userId, createdNGOId)
	if err != nil {
		return "", err
	}
	return createdNGOId, nil
}

func GetNGOsAdmins() (NGOsAdmins []models.NGOAdmin, err error) {
	rows, err := db.Query("SELECT * FROM ngos_admins")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var ngoAdmin models.NGOAdmin
		if err := rows.Scan(&ngoAdmin.UserId, &ngoAdmin.NGOId); err != nil {
			return nil, err
		}
		NGOsAdmins = append(NGOsAdmins, ngoAdmin)
	}
	return NGOsAdmins, nil
}

func GetNGOs() (NGOs []models.NGO, err error) {
	rows, err := db.Query("SELECT * FROM ngos")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var NGO models.NGO
		if err := rows.Scan(&NGO.Id, &NGO.Name, &NGO.Description, &NGO.Rating, &NGO.CreatorId); err != nil {
			return nil, err
		}
		NGOs = append(NGOs, NGO)
	}

	rows.Close()
	return NGOs, nil
}

func GetUserNGO(companyId *string) (ngo models.NGO, err error) {
	if err := db.QueryRow("SELECT * FROM ngos WHERE id = $1", companyId).
		Scan(&ngo.Id, &ngo.Name, &ngo.Description, &ngo.Rating, &ngo.CreatorId); err != nil {
		return ngo, err
	}
	return ngo, err
}
