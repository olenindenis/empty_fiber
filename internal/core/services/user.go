package services

import (
	"database/sql"
	"errors"

	"envs/internal/core/domain"
	"envs/internal/core/ports"
	"envs/internal/dto"
)

type User struct {
	userRepository ports.UserRepository
}

var _ ports.UserService = (*User)(nil)

func NewUser(userRepository ports.UserRepository) *User {
	return &User{userRepository: userRepository}
}

func (us *User) List(filter dto.ListFilter) ([]domain.User, error) {
	users, err := us.userRepository.List(filter)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return []domain.User{}, err
	}

	return users, nil
}

func (us *User) Show(id uint) (domain.User, error) {
	user, err := us.userRepository.Find(id)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return domain.User{}, err
	}

	return user, nil
}

func (us *User) Update(user domain.User) error {
	err := us.userRepository.Update(user)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return err
	}

	return nil
}

func (us *User) Delete(id uint) error {
	err := us.userRepository.Delete(id)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return err
	}

	return nil
}
