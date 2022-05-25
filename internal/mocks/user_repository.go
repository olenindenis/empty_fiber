package mocks

import (
	"envs/internal/core/domain"
	"envs/internal/core/ports"
)

var _ ports.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	UserID   uint
	Password string
}

func NewUserRepository(userID uint) *UserRepository {
	return &UserRepository{UserID: userID}
}

func (ur *UserRepository) Store(name, email, password string) (domain.User, error) {
	return domain.User{
		ID:    ur.UserID,
		Name:  name,
		Email: email,
	}, nil
}

func (ur *UserRepository) FindByEmail(email string) (domain.User, error) {
	return domain.User{ID: ur.UserID, Password: ur.Password}, nil
}

func (ur *UserRepository) Find(id uint) (domain.User, error) {
	return domain.User{ID: id, Password: ur.Password}, nil
}

func (ur *UserRepository) SetPassword(password string) {
	ur.Password = password
}

func (ur *UserRepository) List(limit, offset uint) ([]domain.User, error) {
	return nil, nil
}

func (ur *UserRepository) Update(user domain.User) error {
	return nil
}

func (ur *UserRepository) Delete(id uint) error {
	return nil
}
