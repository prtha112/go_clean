package presenter

import (
	model "go_clean/App/Layer/Entity/Model"
	presenter "go_clean/App/Layer/Usecase/Presenter"
)

type userPresenter struct{}

func NewUserPresenter() presenter.UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(us []*model.User) []*model.User {
	for _, u := range us {
		u.Name = "Mr." + u.Name
	}
	return us
}
