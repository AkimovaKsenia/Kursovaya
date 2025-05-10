package entities

import "github.com/lib/pq"

type Film struct {
	ID            int      `json:"id" db:"id"`
	Name          string   `json:"name" db:"name"`
	Description   string   `json:"description" db:"description"`
	Photo         string   `json:"photo" db:"photo"`
	CastList      []string `json:"cast_list" db:"cast_list"`
	FilmStudioID  int      `json:"film_studio_id" db:"film_studio_id"`
	DurationInMin int      `json:"duration_in_min" db:"duration_in_min"`
}

type FilmStudio struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Director struct {
	ID  int    `json:"id" db:"id"`
	FIO string `json:"fio" db:"fio"`
}

type Operator struct {
	ID  int    `json:"id" db:"id"`
	FIO string `json:"fio" db:"fio"`
}

type Genre struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type CreateFilm struct {
	Name          string         `json:"name" db:"name" form:"name"`
	Description   string         `json:"description" db:"description" form:"description"`
	Photo         string         `json:"photo" db:"photo" form:"photo"`
	CastList      pq.StringArray `json:"cast_list" db:"cast_list" form:"cast_list" swaggertype:"array,string"`
	FilmStudioID  int            `json:"film_studio_id" db:"film_studio_id" form:"film_studio_id"`
	DurationInMin int            `json:"duration_in_min" db:"duration_in_min" form:"duration_in_min"`
	DirectorIDs   []int          `json:"director_ids" db:"director_ids" form:"director_ids"`
	OperatorIDs   []int          `json:"operator_ids" db:"operator_ids" form:"operator_ids"`
	GenreIDs      []int          `json:"genre_ids" db:"genre_ids" form:"genre_ids"`
}

type FilmFull struct {
	ID             int            `json:"id" db:"id"`
	Name           string         `json:"name" db:"name"`
	Description    string         `json:"description" db:"description"`
	Photo          string         `json:"photo" db:"photo"`
	CastList       pq.StringArray `json:"cast_list" db:"cast_list" swaggertype:"array,string"`
	FilmStudioName string         `json:"film_studio_name" db:"film_studio_name"`
	DurationInMin  int            `json:"duration_in_min" db:"duration_in_min"`
	Directors      pq.StringArray `json:"directors" db:"-" swaggertype:"array,string"`
	Operators      pq.StringArray `json:"operators" db:"-" swaggertype:"array,string"`
	Genres         pq.StringArray `json:"genres" db:"-" swaggertype:"array,string"`
}

type UpdateFilm struct {
	ID            int            `json:"id" db:"id" form:"id"`
	Name          string         `json:"name" db:"name" form:"name"`
	Description   string         `json:"description" db:"description" form:"description"`
	Photo         string         `json:"photo" db:"photo" form:"photo"`
	CastList      pq.StringArray `json:"cast_list" db:"cast_list" form:"cast_list" swaggertype:"array,string"`
	FilmStudioID  int            `json:"film_studio_id" db:"film_studio_id" form:"film_studio_id"`
	DurationInMin int            `json:"duration_in_min" db:"duration_in_min" form:"duration_in_min"`
	DirectorIDs   []int          `json:"director_ids" db:"director_ids" form:"director_ids"`
	OperatorIDs   []int          `json:"operator_ids" db:"operator_ids" form:"operator_ids"`
	GenreIDs      []int          `json:"genre_ids" db:"genre_ids" form:"genre_ids"`
}
