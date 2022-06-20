package entity

type User struct {
	ID int `json:"id" gorm:"primary_key"`

	UserId string `json:"user_id" gorm:"column:user_id"`

	Username string `json:"username" gorm:"uniqueIndex;type:varchar(16);not null"`

	Password string `json:"password" gorm:"type:varchar(255)"`

	Email string `json:"email" gorm:"type:varchar(128);not null"`

	Phone string `json:"phone" gorm:"type:varchar(16)"`

	Role string `json:"role" gorm:"type:varchar(16);not null;default:'user'"`

	CreatedAt int64 `json:"created_at"`

	UpdatedAt int64 `json:"updated_at"`

	DeletedAt int64 `json:"deleted_at"`
}
