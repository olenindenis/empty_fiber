package services

import (
	"envs/internal/core/domain"
	"envs/internal/mocks"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShow(t *testing.T) {
	userID := uint(1)
	service := NewUserService(mocks.NewUserRepository(userID))

	user, err := service.Show(userID)
	assert.Equal(t, nil, err, err)
	assert.Equal(t, userID, user.ID, "userID error")
}

func TestList(t *testing.T) {
	userID := uint(1)
	service := NewUserService(mocks.NewUserRepository(userID))

	_, err := service.List(10, 0)
	assert.Equal(t, nil, err, err)
}

func TestUpdate(t *testing.T) {
	userID := uint(1)
	service := NewUserService(mocks.NewUserRepository(userID))

	user := domain.User{
		ID:    userID,
		Name:  faker.Name(),
		Email: faker.Email(),
	}

	err := service.Update(user)
	assert.Equal(t, nil, err, err)
}

func TestDelete(t *testing.T) {
	userID := uint(1)
	service := NewUserService(mocks.NewUserRepository(userID))

	err := service.Delete(userID)
	assert.Equal(t, nil, err, err)
}
