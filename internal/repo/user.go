package repo

import (
	"context"
	"github.com/rhuandantas/verifymy-test/internal/models"
)

type UserRepo interface {
	Create(ctx context.Context, user models.User) (*models.User, error)
	Update(ctx context.Context, userId int, user models.User) (*models.User, error)
	Delete(ctx context.Context, userId int) (bool, error)
	GetById(ctx context.Context, userId int) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}

type UserRepoImpl struct {
}

func NewUserRepo() UserRepo {
	return &UserRepoImpl{}
}

func (uri *UserRepoImpl) Create(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil
}
func (uri *UserRepoImpl) Update(ctx context.Context, userId int, user models.User) (*models.User, error) {
	return nil, nil
}
func (uri *UserRepoImpl) Delete(ctx context.Context, userId int) (bool, error) {
	return false, nil
}
func (uri *UserRepoImpl) GetById(ctx context.Context, userId int) (*models.User, error) {
	return nil, nil
}
func (uri *UserRepoImpl) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return nil, nil
}
