package entities

type User struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Surname  string `json:"surname" db:"surname"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	RoleID   int    `json:"role_id" db:"role_id"`
}

type GetUser struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Surname string `json:"surname" db:"surname"`
	Email   string `json:"email" db:"email"`
	Role    string `json:"role" db:"role"`
}

type GetUserEmail struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Surname  string `json:"surname" db:"surname"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
}

type CreateUser struct {
	Name     string `json:"name" db:"name"`
	Surname  string `json:"surname" db:"surname"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	RoleID   int    `json:"role_id" db:"role_id"`
}

type LoginUserRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type LoginUserResponse struct {
	Token string `json:"token" db:"token"`
	Role  string `json:"role" db:"role"`
}

type UserRole struct {
	ID   int    `json:"id" db:"id"`
	Role string `json:"role" db:"role"`
}
