package postgres

const (
	createTableRoles = `CREATE TABLE IF NOT EXISTS roles (
	    id SERIAL PRIMARY KEY,
	    name VARCHAR(20) NOT NULL
	)`

	createTableUsers = `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR NOT NULL,
		password VARCHAR NOT NULL,
		name VARCHAR,
		surname VARCHAR,
		role_id INT REFERENCES roles(id) ON DELETE CASCADE
	);`

	createTableCinemaConditions = `CREATE TABLE IF NOT EXISTS cinema_conditions (
		id SERIAL PRIMARY KEY,
		name VARCHAR NOT NULL
	);`

	createTableCinemaCategories = `CREATE TABLE IF NOT EXISTS cinema_categories (
		id SERIAL PRIMARY KEY,
		name VARCHAR NOT NULL
	);`

	createTableCinemas = `CREATE TABLE IF NOT EXISTS cinemas (
		id SERIAL PRIMARY KEY,
		name VARCHAR NOT NULL,
		description TEXT,
		photo TEXT,
		address VARCHAR NOT NULL,
		email VARCHAR,
		phone VARCHAR,
		condition_id INT REFERENCES cinema_conditions(id) ON DELETE SET NULL,
		category_id INT REFERENCES cinema_categories(id) ON DELETE SET NULL
	);`

	createTableCinemaHallTypes = `CREATE TABLE IF NOT EXISTS cinema_hall_types (
		id SERIAL PRIMARY KEY,
		name VARCHAR NOT NULL
	);`

	createTableCinemaHalls = `CREATE TABLE IF NOT EXISTS cinema_halls (
		id SERIAL PRIMARY KEY,
		name VARCHAR NOT NULL,
		capacity INT NOT NULL,
		type_id INT REFERENCES cinema_hall_types(id) ON DELETE SET NULL,
		cinema_id INT REFERENCES cinemas(id) ON DELETE CASCADE
	);`

	createTableFilmStudios = `CREATE TABLE IF NOT EXISTS film_studios (
		id SERIAL PRIMARY KEY,
		name VARCHAR NOT NULL
	);`

	createTableFilms = `CREATE TABLE IF NOT EXISTS films (
		id SERIAL PRIMARY KEY,
		name VARCHAR NOT NULL,
		description TEXT,
		photo TEXT,
		cast_list VARCHAR[],
		film_studio_id INT REFERENCES film_studios(id) ON DELETE SET NULL,
		duration_in_min INT NOT NULL
	);`

	createTableDirectors = `CREATE TABLE IF NOT EXISTS directors (
		id SERIAL PRIMARY KEY,
		fio VARCHAR NOT NULL
	);`

	createTableFilmsDirectors = `CREATE TABLE IF NOT EXISTS films_directors (
		film_id INT REFERENCES films(id) ON DELETE CASCADE,
		director_id INT REFERENCES directors(id) ON DELETE CASCADE,
		PRIMARY KEY (film_id, director_id)
	);`

	createTableOperators = `CREATE TABLE IF NOT EXISTS operators (
		id SERIAL PRIMARY KEY,
		fio VARCHAR NOT NULL
	);`

	createTableFilmsOperators = `CREATE TABLE IF NOT EXISTS films_operators (
		film_id INT REFERENCES films(id) ON DELETE CASCADE,
		operator_id INT REFERENCES operators(id) ON DELETE CASCADE,
		PRIMARY KEY (film_id, operator_id)
	);`

	createTableGenres = `CREATE TABLE IF NOT EXISTS genres (
		id SERIAL PRIMARY KEY,
		name VARCHAR NOT NULL
	);`

	createTableFilmsGenres = `CREATE TABLE IF NOT EXISTS films_genres (
		film_id INT REFERENCES films(id) ON DELETE CASCADE,
		genre_id INT REFERENCES genres(id) ON DELETE CASCADE,
		PRIMARY KEY (film_id, genre_id)
	);`

	createTableCinemaSessions = `CREATE TABLE IF NOT EXISTS cinema_sessions (
		id SERIAL PRIMARY KEY,
		cinema_id INT REFERENCES cinemas(id) ON DELETE CASCADE,
		film_id INT REFERENCES films(id) ON DELETE CASCADE,
		date DATE NOT NULL,
		time TIME NOT NULL,
		cinema_hall_id INT REFERENCES cinema_halls(id) ON DELETE CASCADE,
		price DECIMAL(10,2) NOT NULL
	); `
)
