package model

import "github.com/jinzhu/gorm"

type User struct {
	*Model
	User     string `json:"user"`
	State    uint8  `json:"state"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ParentID string `json:"parent_id"`
}

func (u User) TableName() string {
	return "user"
}

func (u User) Count(db *gorm.DB) (int, error) {
	var count int
	if u.User != "" {
		db = db.Where("name = ?", u.User)
	}
	db = db.Where("state = ?", u.State)
	if err := db.Model(&u).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (u User) List(db *gorm.DB, pageOffset, pageSize int) ([]*User, error) {
	var Users []*User
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if u.User != "" {
		db = db.Where("name = ?", u.User)
	}
	db = db.Where("state = ?", u.State)
	if err = db.Where("is_del = ?", 0).Find(&Users).Error; err != nil {
		return nil, err
	}

	return Users, nil
}

func (u User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u User) Update(db *gorm.DB, value interface{}) error {
	return db.Model(&User{}).Where("id = ? AND is_del = ?", u.ID, 0).Update(value).Error
}

func (u User) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", u.Model.ID, 0).Delete(&u).Error
}

func (u User) CheckEmail(db *gorm.DB) (bool, error) {
	var count int
	err := db.Model(&u).Where("email = ?", u.Email).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (u User) CheckUser(db *gorm.DB) (bool, error) {
	var count int
	err := db.Model(&u).Where("user = ?", u.User).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (u User) Get(db *gorm.DB) (User, error) {
	var auth User
	if u.User != "" {
		db = db.Where("user = ? AND password = ? AND is_del = ?", u.User, u.Password, 0)
	} else if u.Email != "" && u.User == "" {
		db = db.Where("email = ? AND password = ? AND is_del = ?", u.Email, u.Password, 0)
	}
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}
