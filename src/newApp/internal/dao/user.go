package dao

import (
	"NewApp/internal/model"
	"NewApp/pkg/app"
)

func (d *Dao) CountUser(name string, state uint8) (int, error) {
	User := model.User{Name: name, State: state}
	return User.Count(d.engine)
}

func (d *Dao) GetUserList(name string, state uint8, page, pageSize int) ([]*model.User, error) {
	User := model.User{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return User.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateUser(name string, state uint8, createdBy string) error {
	User := model.User{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}

	return User.Create(d.engine)
}

func (d *Dao) UpdateUser(id uint32, name string, state uint8, modifiedBy string) error {
	User := model.User{
		Model: &model.Model{ID: id},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	return User.Update(d.engine, values)
}

func (d *Dao) DeleteUser(id uint32) error {
	User := model.User{Model: &model.Model{ID: id}}
	return User.Delete(d.engine)
}
