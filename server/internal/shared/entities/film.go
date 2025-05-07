package entities

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
