package model

import "github.com/jinzhu/gorm"

type User struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t User) TableName() string {
	return "user"
}

func (t User) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t User) List(db *gorm.DB, pageOffset, pageSize int) ([]*User, error) {
	var Users []*User
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&Users).Error; err != nil {
		return nil, err
	}

	return Users, nil
}

func (t User) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t User) Update(db *gorm.DB, value interface{}) error {
	return db.Model(&User{}).Where("id = ? AND is_del = ?", t.ID, 0).Update(value).Error
}

func (t User) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}
