package postgres

import (
	"fmt"
	"kino/internal/shared/config"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	DB *sqlx.DB
}

func NewDatabase(envConf *config.Config) *DB {
	connectionString := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		envConf.Db.User,
		envConf.Db.Password,
		envConf.Db.Host,
		envConf.Db.Port,
		envConf.Db.Database,
	)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatalf("error opening database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("error pinging database connection: %v", err)
	}

	CreateTable(db)

	return &DB{db}
}

func CreateTable(db *sqlx.DB) {
	db.MustExec(createTableRoles)
	db.MustExec(createTableUsers)
	db.MustExec(createTableCinemaConditions)
	db.MustExec(createTableCinemaCategories)
	db.MustExec(createTableCinemas)
	db.MustExec(createTableCinemaHallTypes)
	db.MustExec(createTableCinemaHalls)
	db.MustExec(createTableFilmStudios)
	db.MustExec(createTableFilms)
	db.MustExec(createTableDirectors)
	db.MustExec(createTableFilmsDirectors)
	db.MustExec(createTableOperators)
	db.MustExec(createTableFilmsOperators)
	db.MustExec(createTableGenres)
	db.MustExec(createTableFilmsGenres)
	db.MustExec(createTableCinemaSessions)
}
