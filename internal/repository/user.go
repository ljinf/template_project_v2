package repository

import (
	"context"
	"github.com/ljinf/template_project_v2/internal/model"
)

type UserRepository interface {
	Cache
	SelectById(ctx context.Context, uid string) (*model.User, error)
}

type userRepository struct {
	*Repository
}

func (u *userRepository) SelectById(ctx context.Context, uid string) (*model.User, error) {
	user := &model.User{
		ID: 123,
	}
	return user, nil
}

func NewUserRepository(r *Repository) UserRepository {
	return &userRepository{
		Repository: r,
	}
}
