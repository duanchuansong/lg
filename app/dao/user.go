package dao

import "zj/app/model"

func (d *Dao) Login(u *model.User) (*model.User, error) {
	err := d.db.Where("user_name=?", u.UserName).Take(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (d *Dao) GetUserByName(name string) (u *model.User, err error) {
	err = d.db.Where("user_name=?", name).Take(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (d *Dao) CreateUser(user *model.User) error {
	return d.db.Create(user).Error
}
