package dbm

import "github.com/jinzhu/gorm"

type UserAuth struct {
	gorm.Model
	// access time
	AccessTime string `json:"access_time,omitempty"`

	// access token
	AccessToken string `json:"access_token,omitempty"`

	// 用户名称
	AuthCode string `json:"auth_code,omitempty"`

	// 用户密码
	AuthType string `json:"auth_type,omitempty"`

	// 额外信息，直接存JSON，比如三方用户信息
	Extend string `json:"extend,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`

	// user id
	UserID int64 `json:"user_id,omitempty"`

	// 三方用户ID
	AuthId string `json:"auth_id,omitempty"`
}

func (u UserAuth) Table() string {
	return "user_auth"
}

func (r *Manager) InserUserAuth(user UserAuth)  {
	r.DB.Create(user)
}

func (r *Manager) UpdateUserAuthUserId(authId string, userId int64)  {
	r.DB.Table("user_auth").Where("auth_id=?",authId).Update(struct {
		UserID int64
	}{
		UserID: userId,
	})
}

