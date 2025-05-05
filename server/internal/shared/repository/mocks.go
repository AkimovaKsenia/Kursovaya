package repository

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (r *Repository) CreateMocks() {
	r.CreateRoles()
	r.CreateUsers()
}

func (r *Repository) CreateRoles() {
	var count int
	err := r.DB.DB.QueryRow("SELECT COUNT(*) FROM roles").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		for _, role := range mockRoles {
			_, err = r.DB.CreateRole(&role)
			if err != nil {
				log.Fatal(fmt.Errorf("error creating mock roles: %w", err))
			}
		}

		log.Println("successful creating mock roles")
	} else {
		log.Println("roles already exist, skipping creation")
	}
}

func (r *Repository) CreateUsers() {
	var count int
	err := r.DB.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		for _, u := range mockUsers {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
			if err != nil {
				log.Fatal(err)
			}
			u.Password = string(hashedPassword)

			_, err = r.DB.CreateUser(&u)
			if err != nil {
				log.Fatal(fmt.Errorf("error creating mock users: %w", err))
			}
		}

		log.Println("successful creating mock users")
	} else {
		log.Println("users already exist, skipping creation")
	}
}
