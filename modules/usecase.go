package modules

import (
	"context"
	"tutorial/entity"
)

type UserUseCase struct {
	Repo Repository
}

type UseCase interface {
	GetByID(ctx context.Context, id int32) (*entity.User, error)
	GetByName(ctx context.Context, name string) ([]*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id int32) error
}

func NewUseCase(repo Repository) UseCase {
	return &UserUseCase{
		Repo: repo,
	}
}

func (u UserUseCase) GetByID(ctx context.Context, id int32) (*entity.User, error) {
	res, err := u.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u UserUseCase) GetByName(ctx context.Context, name string) ([]*entity.User, error) {
	res, err := u.Repo.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u UserUseCase) CreateUser(ctx context.Context, user *entity.User) error {
	err := u.Repo.CreateUser(ctx, user.Name, user.Email, user.Age)
	if err != nil {
		return err
	}
	return nil
}
func (u UserUseCase) UpdateUser(ctx context.Context, user *entity.User) error {
	err := u.Repo.UpdateUser(ctx, user.Name, user.Email, user.Age, user.User_id)
	if err != nil {
		return err
	}
	return nil
}
func (u UserUseCase) DeleteUser(ctx context.Context, id int32) error {
	err := u.Repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
