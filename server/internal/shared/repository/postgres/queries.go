package postgres

const (
	createTableRoles = `CREATE TABLE IF NOT EXISTS roles (
	    id SERIAL PRIMARY KEY,
	    name VARCHAR(20) NOT NULL
	)`

	createTableUsers = `
		CREATE TABLE IF NOT EXISTS users (
		  id SERIAL PRIMARY KEY,
		  name VARCHAR,
		  surname VARCHAR,
		  role_id INT REFERENCES roles(id) ON DELETE CASCADE,
		  email VARCHAR,
		  password VARCHAR
	);
	`
)
