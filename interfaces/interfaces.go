package interfaces

// RegistReq struct for register request
type RegistReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	LevelID   int    `json:"level_id"`
}

// LoginReq struct for login request
type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User table
type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	ProfileImage string `json:"profile_image"`
	PhoneNumber  string `json:"phone_number"`
	LevelID      int    `json:"level_id"`
}
