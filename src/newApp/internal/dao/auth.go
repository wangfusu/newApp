package dao

import "NewApp/internal/model"

func (d *Dao) GetAuth(user, password string) (model.Auth, error) {
	auth := model.Auth{User: user, Password: password}
	return auth.Get(d.engine)
}
