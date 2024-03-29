package Repository

import (
	"errors"
	model "go_clean/App/Layer/Entity/Model"
	repository "go_clean/App/Layer/Usecase/Repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindAll(u []*model.User) ([]*model.User, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *userRepository) Create(u *model.User) (*model.User, error) {
	if err := ur.db.Create(u).Error; !errors.Is(err, nil) {
		return nil, err
	}

	return u, nil
}
