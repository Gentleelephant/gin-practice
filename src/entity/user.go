package entity

type User struct {
	ID int `json:"id" gorm:"primary_key"`

	Username string `json:"username"`

	Password string `json:"password"`

	Email string `json:"email"`

	Phone string `json:"phone"`

	CreatedAt int64 `json:"created_at"`

	UpdatedAt int64 `json:"updated_at"`

	DeletedAt int64 `json:"deleted_at"`
}
