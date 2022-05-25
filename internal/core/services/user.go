package services

import (
	"database/sql"
	"envs/internal/core/domain"
	"envs/internal/core/ports"
	"errors"
)

type UserService struct {
	userRepository ports.UserRepository
}

var _ ports.UserService = (*UserService)(nil)

func NewUserService(userRepository ports.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (us *UserService) List(limit, offset uint) ([]domain.User, error) {
	users, err := us.userRepository.List(limit, offset)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return []domain.User{}, err
	}

	return users, nil
}

func (us *UserService) Show(id uint) (domain.User, error) {
	user, err := us.userRepository.Find(id)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return domain.User{}, err
	}

	return user, nil
}

func (us *UserService) Update(user domain.User) error {
	err := us.userRepository.Update(user)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return err
	}

	return nil
}

func (us *UserService) Delete(id uint) error {
	err := us.userRepository.Delete(id)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return err
	}

	return nil
}
