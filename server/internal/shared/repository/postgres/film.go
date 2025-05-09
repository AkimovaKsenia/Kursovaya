package postgres

import (
	"fmt"
	"kino/internal/shared/entities"
)

func (db *DB) CreateFilmStudio(fs *entities.FilmStudio) (int, error) {
	query := `
        INSERT INTO film_studios (name)
        VALUES (:name)
        RETURNING id
    `

	var id int
	stmt, err := db.DB.PrepareNamed(query)
	if stmt == nil {
		return 0, fmt.Errorf("create film studio error preparing statement: %w", err)
	}
	err = stmt.Get(&id, fs)
	if err != nil {
		return 0, fmt.Errorf("error creating film studio: %w", err)
	}

	return id, nil
}

func (db *DB) CreateGenre(g *entities.Genre) (int, error) {
	query := `
        INSERT INTO genres (name)
        VALUES (:name)
        RETURNING id
    `

	var id int
	stmt, err := db.DB.PrepareNamed(query)
	if stmt == nil {
		return 0, fmt.Errorf("create genre error preparing statement: %w", err)
	}
	err = stmt.Get(&id, g)
	if err != nil {
		return 0, fmt.Errorf("error creating genre: %w", err)
	}

	return id, nil
}

func (db *DB) CreateOperator(o *entities.Operator) (int, error) {
	query := `
        INSERT INTO operators (fio)
        VALUES (:fio)
        RETURNING id
    `

	var id int
	stmt, err := db.DB.PrepareNamed(query)
	if stmt == nil {
		return 0, fmt.Errorf("create operator error preparing statement: %w", err)
	}
	err = stmt.Get(&id, o)
	if err != nil {
		return 0, fmt.Errorf("error creating operator: %w", err)
	}

	return id, nil
}

func (db *DB) CreateDirector(d *entities.Director) (int, error) {
	query := `
        INSERT INTO directors (fio)
        VALUES (:fio)
        RETURNING id
    `

	var id int
	stmt, err := db.DB.PrepareNamed(query)
	if stmt == nil {
		return 0, fmt.Errorf("create director error preparing statement: %w", err)
	}
	err = stmt.Get(&id, d)
	if err != nil {
		return 0, fmt.Errorf("error creating director: %w", err)
	}

	return id, nil
}

func (db *DB) CreateFilm(f *entities.CreateFilm) (int, error) {
	// Начинаем транзакцию
	tx, err := db.DB.Beginx()
	if err != nil {
		return 0, fmt.Errorf("error starting transaction: %w", err)
	}

	// Откатываем транзакцию в случае ошибки
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 1. Создаем основную запись фильма
	var filmID int
	err = tx.QueryRowx(`
        INSERT INTO films (
            name, 
            description, 
            photo, 
            cast_list, 
            film_studio_id, 
            duration_in_min
        )
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `, f.Name, f.Description, f.Photo, f.CastList, f.FilmStudioID, f.DurationInMin).Scan(&filmID)

	if err != nil {
		return 0, fmt.Errorf("error creating film: %w", err)
	}

	// 2. Добавляем связи с режиссерами
	for _, directorID := range f.DirectorIDs {
		_, err = tx.Exec(`
            INSERT INTO films_directors (film_id, director_id) 
            VALUES ($1, $2)
        `, filmID, directorID)
		if err != nil {
			return 0, fmt.Errorf("error creating film-director relation: %w", err)
		}
	}

	// 3. Добавляем связи с операторами
	for _, operatorID := range f.OperatorIDs {
		_, err = tx.Exec(`
            INSERT INTO films_operators (film_id, operator_id) 
            VALUES ($1, $2)
        `, filmID, operatorID)
		if err != nil {
			return 0, fmt.Errorf("error creating film-operator relation: %w", err)
		}
	}

	// 4. Добавляем связи с жанрами
	for _, genreID := range f.GenreIDs {
		_, err = tx.Exec(`
            INSERT INTO films_genres (film_id, genre_id) 
            VALUES ($1, $2)
        `, filmID, genreID)
		if err != nil {
			return 0, fmt.Errorf("error creating film-genre relation: %w", err)
		}
	}

	// Фиксируем транзакцию
	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("error committing transaction: %w", err)
	}

	return filmID, nil
}
