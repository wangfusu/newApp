package dao

import (
	"NewApp/internal/model"
	"NewApp/pkg/app"
	"NewApp/utils"
	"errors"
)

func (d *Dao) CountUser(name string, state uint8) (int, error) {
	User := model.User{User: name, State: state}
	return User.Count(d.engine)
}

func (d *Dao) GetUserList(name string, state uint8, page, pageSize int) ([]*model.User, error) {
	User := model.User{User: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return User.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateUser(user, password, email string) error {
	id := utils.CreateUUID()
	User := model.User{
		Model:    &model.Model{ID: id},
		User:     user,
		State:    1,
		Email:    email,
		Password: password,
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

func (d *Dao) CheckEmail(email string) error {
	user := model.User{Email: email}
	find, err := user.CheckEmail(d.engine)
	if err != nil {
		return err
	}
	if find {
		return errors.New("该邮箱已被注册")
	} else {
		return nil
	}
}

func (d *Dao) CheckUser(user string) error {
	userTab := model.User{Email: user}
	find, err := userTab.CheckUser(d.engine)
	if err != nil {
		return err
	}
	if find {
		return errors.New("该用户名已被注册")
	} else {
		return nil
	}
}
