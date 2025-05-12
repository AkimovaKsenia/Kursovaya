package entities

type CinemaCondition struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type CinemaCategory struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Cinema struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Photo       string `json:"photo" db:"photo"`
	Address     string `json:"address" db:"address"`
	Email       string `json:"email" db:"email"`
	Phone       string `json:"phone" db:"phone"`
	ConditionID int    `json:"condition_id" db:"condition_id"`
	CategoryID  int    `json:"category_id" db:"category_id"`
}

type CreateCinema struct {
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Photo       string `json:"photo" db:"photo"`
	Address     string `json:"address" db:"address"`
	Email       string `json:"email" db:"email"`
	Phone       string `json:"phone" db:"phone"`
	ConditionID int    `json:"condition_id" db:"condition_id"`
	CategoryID  int    `json:"category_id" db:"category_id"`
}

type GetCinema struct {
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Photo       string `json:"photo" db:"photo"`
	Address     string `json:"address" db:"address"`
	Email       string `json:"email" db:"email"`
	Phone       string `json:"phone" db:"phone"`
	Condition   string `json:"condition" db:"condition"`
	Category    string `json:"category" db:"category"`
}

type CinemaNameAddress struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
}

type CinemaHallType struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type CinemaHall struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Capacity int    `json:"capacity" db:"capacity"`
	TypeID   int    `json:"type_id" db:"type_id"`
	CinemaID int    `json:"cinema_id" db:"cinema_id"`
}

type GetCinemaHall struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Capacity int    `json:"capacity" db:"capacity"`
	Type     string `json:"type" db:"type"`
}
