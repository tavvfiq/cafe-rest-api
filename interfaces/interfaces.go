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

// Menu menu returning struct
type Menu struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Image     string `json:"image_path"`
	Price     string `json:"price"`
	Quantity  string `json:"quantity"`
	Category  string `json:"category"`
	AddedAt   string `json:"added_at"`
	UpdatedAt string `json:"updated_at"`
}

// PageInfo page info for pagination
type PageInfo struct {
	PrevPage string `json:"prevPage"`
	Page     int    `json:"page"`
	NextPage string `json:"nextPage"`
}
