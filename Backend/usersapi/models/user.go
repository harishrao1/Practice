package models

import (
	"database/sql"
	"errors"
	"time"
)

type User struct {
	ID           int     `json:"id"`
	Username     string  `json:"username"`
	Email        string  `json:"email"`
	PasswordHash string  `json:"-"` // omit in JSON responses
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	PhoneNumber  string  `json:"phone_number"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

// ✅ Create User
func (u *User) Create(db *sql.DB) error {
	query := `
		INSERT INTO users (username, email, password_hash, first_name, last_name, phone_number, latitude, longitude, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
	`
	result, err := db.Exec(query, u.Username, u.Email, u.PasswordHash, u.FirstName, u.LastName, u.PhoneNumber, u.Latitude, u.Longitude)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = int(id)
	u.CreatedAt = time.Now().Format(time.RFC3339)
	u.UpdatedAt = time.Now().Format(time.RFC3339)

	return nil
}

// ✅ Get All Users
func GetAllUsers(db *sql.DB) ([]User, error) {
	query := `
		SELECT id, username, email, first_name, last_name, phone_number, latitude, longitude, created_at, updated_at
		FROM users
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		if err := rows.Scan(
			&u.ID,
			&u.Username,
			&u.Email,
			&u.FirstName,
			&u.LastName,
			&u.PhoneNumber,
			&u.Latitude,
			&u.Longitude,
			&u.CreatedAt,
			&u.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

// ✅ Get User By ID
func GetUserByID(db *sql.DB, id int) (*User, error) {
	query := `
		SELECT id, username, email, first_name, last_name, phone_number, latitude, longitude, created_at, updated_at
		FROM users WHERE id = ?
	`
	var u User
	err := db.QueryRow(query, id).Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.PhoneNumber,
		&u.Latitude,
		&u.Longitude,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &u, nil
}

// ✅ Update User
func (u *User) Update(db *sql.DB) error {
	if u.ID == 0 {
		return errors.New("missing user ID")
	}

	query := `
		UPDATE users
		SET username = ?, email = ?, first_name = ?, last_name = ?, phone_number = ?, latitude = ?, longitude = ?, updated_at = NOW()
		WHERE id = ?
	`

	_, err := db.Exec(query, u.Username, u.Email, u.FirstName, u.LastName, u.PhoneNumber, u.Latitude, u.Longitude, u.ID)
	if err != nil {
		return err
	}

	u.UpdatedAt = time.Now().Format(time.RFC3339)
	return nil
}

// ✅ Delete User
func DeleteUser(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = ?`

	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
