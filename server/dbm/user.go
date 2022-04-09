package dbm

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	// avtor
	Avtor string `json:"avtor,omitempty"`
	// email
	Email string `json:"email,omitempty"`
	// login time
	LoginTime int64 `json:"login_time,omitempty"`
	// 用户名称
	Name string `json:"name,omitempty"`
	// 用户密码
	Password string `json:"password,omitempty"`
	// phone
	Phone string `json:"phone,omitempty"`
	// register time
	RegisterTime int64 `json:"register_time,omitempty"`
	// status
	Status int64 `json:"status,omitempty"`
}

func (u User) Table() string {
	return "user"
}

func (r *Manager) Inser(user User)  {
	r.DB.Create(user)
}
