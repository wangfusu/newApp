package model

import "github.com/jinzhu/gorm"

type Auth struct {
	ID       uint32 `gorm:"primary_key" json:"id"`
	User     string `json:"user"`
	Password string `json:"password"`
	IsDel    uint8  `json:"is_del"`
}

func (a Auth) TableName() string {
	return "user_auth"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("user = ? AND password = ? AND is_del = ?", a.User, a.Password, 0)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}
