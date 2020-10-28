package mysql

// RegistReq struct for register request
type RegistReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	LevelID   int    `json:"level_id"`
}

// User table
type User struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	ProfileImage string `json:"profile_image"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
	LevelID      int    `json:"level_id"`
}
