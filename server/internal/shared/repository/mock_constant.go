package repository

import "kino/internal/shared/entities"

var (
	mockRoles = []entities.Role{
		{
			Name: "worker",
		},
		{
			Name: "admin",
		},
	}

	mockUsers = []entities.User{
		{
			Name:     "Данек",
			Surname:  "Гевинов",
			Email:    "worker@yandex.ru",
			Password: "worker",
			RoleID:   1,
		},
		{
			Name:     "Ксесша",
			Surname:  "Акимова",
			Email:    "admin@yandex.ru",
			Password: "admin",
			RoleID:   2,
		},
	}
)
