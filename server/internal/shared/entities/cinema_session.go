package entities

type CinemaSession struct {
	ID           int     `json:"id" db:"id"`
	CinemaID     int     `json:"cinema_id" db:"cinema_id"`
	FilmID       int     `json:"film_id" db:"film_id"`
	Date         string  `json:"date" db:"date"`
	Time         string  `json:"time" db:"time"`
	CinemaHallID int     `json:"cinema_hall_id" db:"cinema_hall_id"`
	Price        float64 `json:"price" db:"price"`
}
