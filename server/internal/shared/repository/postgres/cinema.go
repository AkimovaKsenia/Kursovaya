package postgres

import (
	"fmt"
	"kino/internal/shared/entities"
)

func (db *DB) CreateCinemaCondition(condition *entities.CinemaCondition) (int, error) {
	var id int
	err := db.DB.QueryRowx(`
        INSERT INTO cinema_conditions (name)
        VALUES ($1)
        RETURNING id
    `, condition.Name).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error creating cinema condition: %w", err)
	}
	return id, nil
}

func (db *DB) GetAllCinemaConditions() ([]entities.CinemaCondition, error) {
	var conditions []entities.CinemaCondition
	err := db.DB.Select(&conditions, `
        SELECT id, name 
        FROM cinema_conditions 
        ORDER BY name
    `)
	if err != nil {
		return nil, fmt.Errorf("error getting cinema conditions: %w", err)
	}
	return conditions, nil
}

func (db *DB) CreateCinemaCategory(category *entities.CinemaCategory) (int, error) {
	var id int
	err := db.DB.QueryRowx(`
        INSERT INTO cinema_categories (name)
        VALUES ($1)
        RETURNING id
    `, category.Name).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error creating cinema category: %w", err)
	}
	return id, nil
}

func (db *DB) GetAllCinemaCategories() ([]entities.CinemaCategory, error) {
	var categories []entities.CinemaCategory
	err := db.DB.Select(&categories, `
        SELECT id, name 
        FROM cinema_categories 
        ORDER BY name
    `)
	if err != nil {
		return nil, fmt.Errorf("error getting cinema categories: %w", err)
	}
	return categories, nil
}

func (db *DB) CreateCinema(c *entities.CreateCinema) (int, error) {
	var id int
	err := db.DB.QueryRowx(`
        INSERT INTO cinemas (
            name, description, photo, address, 
            email, phone, condition_id, category_id
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id
    `, c.Name, c.Description, c.Photo, c.Address,
		c.Email, c.Phone, c.ConditionID, c.CategoryID).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error creating cinema: %w", err)
	}
	return id, nil
}

func (db *DB) GetAllCinemasAddressName() ([]entities.CinemaAddressName, error) {
	var cinemas []entities.CinemaAddressName
	err := db.DB.Select(&cinemas, `
        SELECT id, name, address 
        FROM cinemas 
        ORDER BY name
    `)
	if err != nil {
		return nil, fmt.Errorf("error getting cinemas address and name: %w", err)
	}
	return cinemas, nil
}

func (db *DB) CreateCinemaHallType(hallType *entities.CinemaHallType) (int, error) {
	var id int
	err := db.DB.QueryRowx(`
        INSERT INTO cinema_hall_types (name)
        VALUES ($1)
        RETURNING id
    `, hallType.Name).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error creating cinema hall type: %w", err)
	}
	return id, nil
}

func (db *DB) GetAllCinemaHallTypes() ([]entities.CinemaHallType, error) {
	var hallTypes []entities.CinemaHallType
	err := db.DB.Select(&hallTypes, `
        SELECT id, name 
        FROM cinema_hall_types 
        ORDER BY name
    `)
	if err != nil {
		return nil, fmt.Errorf("error getting cinema hall types: %w", err)
	}
	return hallTypes, nil
}

func (db *DB) CreateCinemaHall(hall *entities.CinemaHall) (int, error) {
	var id int
	err := db.DB.QueryRowx(`
        INSERT INTO cinema_halls (
            name, capacity, type_id, cinema_id
        )
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `, hall.Name, hall.Capacity, hall.TypeID, hall.CinemaID).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error creating cinema hall: %w", err)
	}
	return id, nil
}

func (db *DB) GetAllCinemaHallsByID(cinemaID int) ([]entities.GetCinemaHall, error) {
	var halls []entities.GetCinemaHall
	err := db.DB.Select(&halls, `
        SELECT 
            ch.id, 
            ch.name, 
            ch.capacity, 
            cht.name as type
        FROM cinema_halls ch
        LEFT JOIN cinema_hall_types cht ON ch.type_id = cht.id
        WHERE ch.cinema_id = $1
        ORDER BY ch.name
    `, cinemaID)
	if err != nil {
		return nil, fmt.Errorf("error getting cinema halls: %w", err)
	}
	return halls, nil
}

func (db *DB) GetCinemaHallByID(id int) (*entities.GetCinemaHall, error) {
	var hall entities.GetCinemaHall
	err := db.DB.Get(&hall, `
        SELECT 
            ch.id, 
            ch.name, 
            ch.capacity, 
            cht.name as type
        FROM cinema_halls ch
        LEFT JOIN cinema_hall_types cht ON ch.type_id = cht.id
        WHERE ch.id = $1
        ORDER BY ch.name
    `, id)
	if err != nil {
		return nil, fmt.Errorf("error getting cinema halls: %w", err)
	}
	return &hall, nil
}

func (db *DB) GetCinemaByID(id int) (*entities.GetCinema, error) {
	var cinema entities.GetCinema
	err := db.DB.Get(&cinema, `
        SELECT 
            c.name,
            c.description,
            c.photo,
            c.address,
            c.email,
            c.phone,
            cc.name as condition,
            ccat.name as category
        FROM cinemas c
        LEFT JOIN cinema_conditions cc ON c.condition_id = cc.id
        LEFT JOIN cinema_categories ccat ON c.category_id = ccat.id
        WHERE c.id = $1
    `, id)
	if err != nil {
		return nil, fmt.Errorf("error getting cinema: %w", err)
	}
	return &cinema, nil
}

func (db *DB) GetAllCinemas() ([]entities.GetCinema, error) {
	var cinemas []entities.GetCinema
	err := db.DB.Select(&cinemas, `
        SELECT 
            c.id,
            c.name,
            c.description,
            c.photo,
            c.address,
            c.email,
            c.phone,
            cc.name as condition,
            ccat.name as category
        FROM cinemas c
        LEFT JOIN cinema_conditions cc ON c.condition_id = cc.id
        LEFT JOIN cinema_categories ccat ON c.category_id = ccat.id
        ORDER BY c.id
    `)
	if err != nil {
		return nil, fmt.Errorf("error getting all cinemas: %w", err)
	}
	return cinemas, nil
}

func (db *DB) UpdateCinema(c *entities.Cinema) error {
	result, err := db.DB.Exec(`
        UPDATE cinemas SET
            name = $1,
            description = $2,
            photo = $3,
            address = $4,
            email = $5,
            phone = $6,
            condition_id = $7,
            category_id = $8
        WHERE id = $9
    `, c.Name, c.Description, c.Photo, c.Address,
		c.Email, c.Phone, c.ConditionID, c.CategoryID, c.ID)

	if err != nil {
		return fmt.Errorf("error updating cinema: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("cinema with id %d not found", c.ID)
	}

	return nil
}

func (db *DB) UpdateCinemaHall(hall *entities.CinemaHall) error {
	result, err := db.DB.Exec(`
        UPDATE cinema_halls SET
            name = $1,
            capacity = $2,
            type_id = $3
        WHERE id = $4
    `, hall.Name, hall.Capacity, hall.TypeID, hall.ID)

	if err != nil {
		return fmt.Errorf("error updating cinema hall: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("cinema hall with id %d not found", hall.ID)
	}

	return nil
}

func (db *DB) DeleteCinema(id int) error {
	result, err := db.DB.Exec(`
        DELETE FROM cinemas 
        WHERE id = $1
    `, id)

	if err != nil {
		return fmt.Errorf("error deleting cinema: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("cinema with id %d not found", id)
	}

	return nil
}

func (db *DB) DeleteCinemaHall(id int) error {
	result, err := db.DB.Exec(`
        DELETE FROM cinema_halls 
        WHERE id = $1
    `, id)

	if err != nil {
		return fmt.Errorf("error deleting cinema hall: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("cinema hall with id %d not found", id)
	}

	return nil
}
