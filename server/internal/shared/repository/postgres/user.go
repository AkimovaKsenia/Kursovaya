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

func (db *DB) GetUserByEmail(email string) (*entities.GetUserEmail, error) {
	user := entities.GetUserEmail{}
	query := `SELECT users.id, users.name, users.surname, roles.name AS role, users.email, users.password
		FROM users
		JOIN public.roles ON users.role_id = roles.id
		WHERE email = $1`
	err := db.DB.Get(&user, query, email)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (db *DB) GetUserRoleById(userID int) (*entities.UserRole, error) {
	query := `
		SELECT users.id, roles.name AS role 
		FROM users JOIN roles ON users.role_id = roles.id 
		WHERE users.id = $1
	`

	var us entities.UserRole
	err := db.DB.Get(&us, query, userID)
	if err != nil {
		return nil, err
	}
	return &us, nil
}

func (db *DB) GetAllUsers() ([]entities.GetUser, error) {
	var users []entities.GetUser
	query := `SELECT 
                users.id, 
                users.name, 
                users.surname, 
                CASE 
                    WHEN roles.name = 'worker' THEN 'Работник'
                    WHEN roles.name = 'admin' THEN 'Администратор'
                    ELSE roles.name
                END AS role,
                users.email
              FROM users
              JOIN public.roles ON users.role_id = roles.id`
	err := db.DB.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}
