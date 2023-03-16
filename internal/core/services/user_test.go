package services

import (
	"testing"

	"envs/internal/core/domain"
	"envs/internal/dto"
	"envs/internal/mocks"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestShow(t *testing.T) {
	userID := uint(1)
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockUserRepository(ctrl)
	service := NewUserService(repo)

	userModel := domain.User{
		ID:    userID,
		Name:  faker.Name(),
		Email: faker.Email(),
	}

	repo.
		EXPECT().
		Find(userID).
		Return(userModel, nil)

	userDomain, err := service.Show(userID)
	assert.NoError(t, err)
	assert.Equal(t, userModel, userDomain)
}

func TestList(t *testing.T) {
	var (
		userID1 uint = 1
		userID2 uint = 2
	)

	ctrl := gomock.NewController(t)
	repo := mocks.NewMockUserRepository(ctrl)
	service, listFilter := NewUserService(repo), dto.ListFilter{}

	usersModel := []domain.User{
		{
			ID:    userID1,
			Name:  faker.Name(),
			Email: faker.Email(),
		},
		{
			ID:    userID2,
			Name:  faker.Name(),
			Email: faker.Email(),
		},
	}

	repo.
		EXPECT().
		List(listFilter).
		Return(usersModel, nil)

	usersDomain, err := service.List(listFilter)
	assert.NoError(t, err)
	assert.Equal(t, usersModel, usersDomain)
}

func TestUpdate(t *testing.T) {
	userID := uint(1)
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockUserRepository(ctrl)
	service := NewUserService(repo)

	user := domain.User{
		ID:    userID,
		Name:  faker.Name(),
		Email: faker.Email(),
	}

	repo.
		EXPECT().
		Update(user).
		Return((error)(nil))

	assert.NoError(t, service.Update(user))
}

func TestDelete(t *testing.T) {
	userID := uint(1)
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockUserRepository(ctrl)
	service := NewUserService(repo)

	repo.
		EXPECT().
		Delete(userID).
		Return((error)(nil))

	assert.NoError(t, service.Delete(userID))
}
