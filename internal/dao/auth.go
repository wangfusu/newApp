package dao

import "NewApp/internal/model"

func (d *Dao) GetAuth(user, password, email string) (model.User, error) {
	userAuth := model.User{User: user, Password: password, Email: email}
	return userAuth.Get(d.engine)
}
