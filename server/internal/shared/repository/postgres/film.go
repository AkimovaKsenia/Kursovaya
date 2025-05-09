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

func (db *DB) GetFilmByID(id int) (*entities.FilmFull, error) {
	// Основная информация о фильме
	var film entities.FilmFull
	err := db.DB.Get(&film, `
        SELECT 
            f.id, 
            f.name, 
            f.description, 
            f.photo, 
            f.cast_list, 
            fs.name as film_studio_name,
            f.duration_in_min
        FROM films f
        LEFT JOIN film_studios fs ON f.film_studio_id = fs.id
        WHERE f.id = $1
    `, id)
	if err != nil {
		return nil, fmt.Errorf("error getting film: %w", err)
	}

	// Получаем режиссеров как массив FIO
	var directors []string
	err = db.DB.Select(&directors, `
        SELECT d.fio
        FROM directors d
        JOIN films_directors fd ON d.id = fd.director_id
        WHERE fd.film_id = $1
    `, id)
	if err != nil {
		return nil, fmt.Errorf("error getting directors: %w", err)
	}
	film.Directors = directors

	// Получаем операторов как массив FIO
	var operators []string
	err = db.DB.Select(&operators, `
        SELECT o.fio
        FROM operators o
        JOIN films_operators fo ON o.id = fo.operator_id
        WHERE fo.film_id = $1
    `, id)
	if err != nil {
		return nil, fmt.Errorf("error getting operators: %w", err)
	}
	film.Operators = operators

	// Получаем жанры как массив названий
	var genres []string
	err = db.DB.Select(&genres, `
        SELECT g.name
        FROM genres g
        JOIN films_genres fg ON g.id = fg.genre_id
        WHERE fg.film_id = $1
    `, id)
	if err != nil {
		return nil, fmt.Errorf("error getting genres: %w", err)
	}
	film.Genres = genres

	return &film, nil
}

func (db *DB) GetAllFilms() ([]entities.FilmFull, error) {
	// Сначала получаем основные данные о фильмах
	var films []entities.FilmFull
	err := db.DB.Select(&films, `
        SELECT 
            f.id, 
            f.name, 
            f.description, 
            f.photo, 
            f.cast_list, 
            fs.name as film_studio_name,
            f.duration_in_min
        FROM films f
        LEFT JOIN film_studios fs ON f.film_studio_id = fs.id
        ORDER BY f.id
    `)
	if err != nil {
		return nil, fmt.Errorf("error getting basic film data: %w", err)
	}

	// Затем для каждого фильма получаем связанные данные
	for i := range films {
		filmID := films[i].ID

		// Получаем режиссеров
		var directors []string
		err = db.DB.Select(&directors, `
            SELECT d.fio
            FROM directors d
            JOIN films_directors fd ON d.id = fd.director_id
            WHERE fd.film_id = $1
        `, filmID)
		if err != nil {
			return nil, fmt.Errorf("error getting directors for film %d: %w", filmID, err)
		}
		films[i].Directors = directors

		// Получаем операторов
		var operators []string
		err = db.DB.Select(&operators, `
            SELECT o.fio
            FROM operators o
            JOIN films_operators fo ON o.id = fo.operator_id
            WHERE fo.film_id = $1
        `, filmID)
		if err != nil {
			return nil, fmt.Errorf("error getting operators for film %d: %w", filmID, err)
		}
		films[i].Operators = operators

		// Получаем жанры
		var genres []string
		err = db.DB.Select(&genres, `
            SELECT g.name
            FROM genres g
            JOIN films_genres fg ON g.id = fg.genre_id
            WHERE fg.film_id = $1
        `, filmID)
		if err != nil {
			return nil, fmt.Errorf("error getting genres for film %d: %w", filmID, err)
		}
		films[i].Genres = genres
	}

	return films, nil
}

func (db *DB) GetAllGenres() ([]entities.Genre, error) {
	var genres []entities.Genre
	err := db.DB.Select(&genres, `
        SELECT id, name
        FROM genres
        ORDER BY name
    `)
	if err != nil {
		return nil, fmt.Errorf("error getting all genres: %w", err)
	}
	return genres, nil
}

func (db *DB) GetAllOperators() ([]entities.Operator, error) {
	var operators []entities.Operator
	err := db.DB.Select(&operators, `
        SELECT id, fio
        FROM operators
        ORDER BY fio
    `)
	if err != nil {
		return nil, fmt.Errorf("error getting all operators: %w", err)
	}
	return operators, nil
}

func (db *DB) GetAllDirectors() ([]entities.Director, error) {
	var directors []entities.Director
	err := db.DB.Select(&directors, `
        SELECT id, fio
        FROM directors
        ORDER BY fio
    `)
	if err != nil {
		return nil, fmt.Errorf("error getting all directors: %w", err)
	}
	return directors, nil
}

func (db *DB) GetAllFilmStudios() ([]entities.FilmStudio, error) {
	var studios []entities.FilmStudio
	err := db.DB.Select(&studios, `
        SELECT id, name
        FROM film_studios
        ORDER BY name
    `)
	if err != nil {
		return nil, fmt.Errorf("error getting all film studios: %w", err)
	}
	return studios, nil
}
