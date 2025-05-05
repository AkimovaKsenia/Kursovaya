package postgres

import (
	"fmt"
	"kino/internal/shared/entities"
)

func (db *DB) CreateUser(u *entities.User) (int, error) {
	query := `
        INSERT INTO users (name, surname, role_id, email, password)
        VALUES (:name, :surname, :role_id, :email, :password)
        RETURNING id
    `

	var id int
	stmt, err := db.DB.PrepareNamed(query)
	if stmt == nil {
		return 0, fmt.Errorf("create user error preparing statement: %w", err)
	}
	err = stmt.Get(&id, u)
	if err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}

	return id, nil
}
