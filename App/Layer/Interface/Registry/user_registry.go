package Registry

import (
	controller "go_clean/App/Layer/Interface/Controllers"
	ip "go_clean/App/Layer/Interface/Presenter"
	ir "go_clean/App/Layer/Interface/Repository"
	interactor "go_clean/App/Layer/Usecase/Interactor"
	up "go_clean/App/Layer/Usecase/Presenter"
	ur "go_clean/App/Layer/Usecase/Repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
