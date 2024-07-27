package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/resq/internal/models"
)

func CreateRequest(request *models.CreateRequest) (createdRequestId string, err error) {
	id, err := gonanoid.New()
	if err != nil {
		return "", err
	}
	err = db.QueryRow("INSERT INTO requests (id, ngo_id, material_id) VALUES ($1, $2, $3) RETURNING id",
		id, request.NGOId, request.MaterialId).
		Scan(&createdRequestId)
	if err != nil {
		return "", err
	}
	return createdRequestId, nil
}

func AcceptRequest(acceptRequestBody *models.AcceptRequest) error {
	err := db.QueryRow("UPDATE requests SET status = 'accepted', message = $2 WHERE id = $1",
		acceptRequestBody.RequestId, acceptRequestBody.Message).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetRequests() (requests []models.Request, err error) {
	rows, err := db.Query("SELECT * FROM requests")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var request models.Request
		rows.Scan(&request.Id, &request.NGOId, &request.MaterialId, &request.CreatedAt, &request.Status, &request.Message)
		if err := rows.Err(); err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}
	return requests, nil
}
