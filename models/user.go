package models

type User struct {
	Id        string `json:"id"`
	LoginName string `json:"login_name"`
	Password  string `json:"password,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`
	UpdatedAt int64 `json:"updated_at,omitempty"`
	DeletedAt int64 `json:"deleted_at"`
}
