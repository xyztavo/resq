package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/resq/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.CreateUserBody) (createdUserId string, err error) {
	nanoid, err := gonanoid.New()
	if err != nil {
		return "", err
	}
	defaultRole := "default"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 0)
	if err != nil {
		return "", err
	}
	err = db.QueryRow("INSERT INTO users (id, name, email, role, password) VALUES ($1, $2, $3, $4, $5) RETURNING id", nanoid, user.Name, user.Email, defaultRole, hashedPassword).Scan(&createdUserId)
	if err != nil {
		return "", err
	}
	return createdUserId, nil
}

func GetUsers() (users []models.User, err error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Role, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	rows.Close()
	return users, nil
}

func GetUserById(userId string) (user models.User, err error) {
	if err := db.QueryRow("SELECT * FROM users WHERE id = $1", userId).
		Scan(&user.Id, &user.Name, &user.Role, &user.Email, &user.Password); err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByEmail(email string) (user models.User, err error) {
	if err := db.QueryRow("SELECT * FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Name, &user.Role, &user.Email, &user.Password); err != nil {
		return user, err
	}
	return user, nil
}

func UpdateNGOUserRole(userId string) error {
	_, err := db.Exec("UPDATE users SET role = 'ngo_admin' WHERE id = $1", userId)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCompanyUserRole(userId string) error {
	_, err := db.Exec("UPDATE users SET role = 'company_admin' WHERE id = $1", userId)
	if err != nil {
		return err
	}
	return nil
}
