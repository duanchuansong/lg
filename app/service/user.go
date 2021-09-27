package service

import (
	"context"
	"github.com/duanchuansong/stk/ecode"
	"github.com/duanchuansong/stk/xlog"
	"gorm.io/gorm"
	"zj/app/model"
)

func (s Service) Login(user *model.User) (*model.User, error) {
	u, err := s.Dao.Login(user)
	if err != nil {
		xlog.Errorf("s.Dao.Login(%v),err:%v", user, err)
		return nil, ecode.RequestErr
	}
	return u, nil
}

func (s Service) CreateUser(c context.Context, user *model.User) (*model.User, error) {
	_, err := s.Dao.GetUserByName(user.UserName)
	if err != nil && gorm.ErrRecordNotFound != err {
		xlog.Errorf("s.Dao.GetUserByName(%s),err:%v", user.UserName, err)
		return nil, ecode.RequestErr
	}
	if err != gorm.ErrRecordNotFound {
		return nil, ecode.UserExist
	}
	if err := s.Dao.CreateUser(user); err != nil {
		xlog.Errorf("s.Dao.CreateUser(%v),err:%v", user, err)
		return nil, ecode.RequestErr
	}
	return user, nil
}
