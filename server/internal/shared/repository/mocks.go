package repository

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (r *Repository) CreateMocks() {
	r.CreateRoles()
	r.CreateUsers()
	r.CreateFilmStudios()
	r.CreateGenres()
	r.CreateDirectors()
	r.CreateOperators()
	r.CreateFilms()
	r.CreateCinemaConditions()
	r.CreateCinemaCategories()
	r.CreateCinemaHallTypes()
	r.CreateCinemas()
	r.CreateCinemaHalls()
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

func (r *Repository) CreateFilmStudios() {
	var count int
	err := r.DB.DB.QueryRow("SELECT COUNT(*) FROM film_studios").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		for _, studio := range mockFilmStudios {
			_, err = r.DB.CreateFilmStudio(&studio)
			if err != nil {
				log.Fatal(fmt.Errorf("error creating mock film studios: %w", err))
			}
		}
		log.Println("successful creating mock film studios")
	} else {
		log.Println("film studios already exist, skipping creation")
	}
}

func (r *Repository) CreateGenres() {
	var count int
	err := r.DB.DB.QueryRow("SELECT COUNT(*) FROM genres").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		for _, genre := range mockGenres {
			_, err = r.DB.CreateGenre(&genre)
			if err != nil {
				log.Fatal(fmt.Errorf("error creating mock genres: %w", err))
			}
		}
		log.Println("successful creating mock genres")
	} else {
		log.Println("genres already exist, skipping creation")
	}
}

func (r *Repository) CreateDirectors() {
	var count int
	err := r.DB.DB.QueryRow("SELECT COUNT(*) FROM directors").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		for _, director := range mockDirectors {
			_, err = r.DB.CreateDirector(&director)
			if err != nil {
				log.Fatal(fmt.Errorf("error creating mock directors: %w", err))
			}
		}
		log.Println("successful creating mock directors")
	} else {
		log.Println("directors already exist, skipping creation")
	}
}

func (r *Repository) CreateOperators() {
	var count int
	err := r.DB.DB.QueryRow("SELECT COUNT(*) FROM operators").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		for _, operator := range mockOperators {
			_, err = r.DB.CreateOperator(&operator)
			if err != nil {
				log.Fatal(fmt.Errorf("error creating mock operators: %w", err))
			}
		}
		log.Println("successful creating mock operators")
	} else {
		log.Println("operators already exist, skipping creation")
	}
}

func (r *Repository) CreateFilms() {
	var count int
	err := r.DB.DB.QueryRow("SELECT COUNT(*) FROM films").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		for _, film := range mockFilms {
			_, err := r.DB.CreateFilm(&film)
			if err != nil {
				log.Fatal(fmt.Errorf("error creating mock film: %w", err))
			}
		}
		log.Println("successful creating mock films")
	} else {
		log.Println("film already exist, skipping creation")
	}
}

func (r *Repository) CreateCinemaConditions() {
	var count int
	err := r.DB.DB.QueryRow("SELECT COUNT(*) FROM cinema_conditions").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		for _, cinemaCondition := range mockCinemaConditions {
			_, err := r.DB.CreateCinemaCondition(&cinemaCondition)
			if err != nil {
				log.Fatal(fmt.Errorf("error creating mock cinema conditions: %w", err))
			}
		}
		log.Println("successful creating mock cinema conditions")
	} else {
		log.Println("cinema conditions already exist, skipping creation")
	}
}

func (r *Repository) CreateCinemaCategories() {
	var count int
	err := r.DB.DB.QueryRow("SELECT COUNT(*) FROM cinema_categories").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		for _, cinemaCategories := range mockCinemaCategories {
			_, err := r.DB.CreateCinemaCategory(&cinemaCategories)
			if err != nil {
				log.Fatal(fmt.Errorf("error creating mock cinema categories: %w", err))
			}
		}
		log.Println("successful creating mock cinema categories")
	} else {
		log.Println("cinema categories already exist, skipping creation")
	}
}

func (r *Repository) CreateCinemaHallTypes() {
	var count int
	err := r.DB.DB.QueryRow("SELECT COUNT(*) FROM cinema_hall_types").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		for _, cinemaHallType := range mockCinemaHallTypes {
			_, err := r.DB.CreateCinemaHallType(&cinemaHallType)
			if err != nil {
				log.Fatal(fmt.Errorf("error creating mock cinema hall types: %w", err))
			}
		}
		log.Println("successful creating mock cinema hall types")
	} else {
		log.Println("cinema hall types already exist, skipping creation")
	}
}

func (r *Repository) CreateCinemas() {
	var count int
	err := r.DB.DB.QueryRow("SELECT COUNT(*) FROM cinemas").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		for _, cinema := range mockCinemas {
			_, err := r.DB.CreateCinema(&cinema)
			if err != nil {
				log.Fatal(fmt.Errorf("error creating mock cinema: %w", err))
			}
		}
		log.Println("successful creating mock cinema")
	} else {
		log.Println("cinema already exist, skipping creation")
	}
}

func (r *Repository) CreateCinemaHalls() {
	var count int
	err := r.DB.DB.QueryRow("SELECT COUNT(*) FROM cinema_halls").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		for _, cinemaHall := range mockCinemaHalls {
			_, err := r.DB.CreateCinemaHall(&cinemaHall)
			if err != nil {
				log.Fatal(fmt.Errorf("error creating mock cinema hall: %w", err))
			}
		}
		log.Println("successful creating mock cinema hall")
	} else {
		log.Println("cinema hall already exist, skipping creation")
	}
}
