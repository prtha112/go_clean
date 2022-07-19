package Presenter

import model "go_clean/App/Layer/Entity/Model"

type userPresenter struct {
}

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(us []*model.User) []*model.User {
	for _, u := range us {
		u.Name = "Mr." + u.Name
	}
	return us
}
