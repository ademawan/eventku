package entities

type User struct {
	UserUid  string `json:"uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`

	CreatedAt interface{} `json:"created_at"`
	UpdatedAt interface{} `json:"updated_at"`
	DeletedAt interface{} `json:"deleted_at"`
}
