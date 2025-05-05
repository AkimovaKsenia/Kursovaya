package postgres

import (
	"fmt"
	"kino/internal/shared/entities"
)

func (db *DB) CreateRole(r *entities.Role) (int, error) {
	query := `
        INSERT INTO roles (name)
        VALUES (:name)
        RETURNING id
    `

	var id int
	stmt, err := db.DB.PrepareNamed(query)
	if stmt == nil {
		return 0, fmt.Errorf("create role error preparing statement: %w", err)
	}
	err = stmt.Get(&id, r)
	if err != nil {
		return 0, fmt.Errorf("error creating role: %w", err)
	}

	return id, nil
}
