package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/rhuandantas/verifymy-test/internal/log"
	"github.com/rhuandantas/verifymy-test/internal/models"
)

//go:generate mockgen -source=$GOFILE -package=mock_repo -destination=../../test/mock/repo/$GOFILE
var (
	RecordNotFoundErr = errors.New("record not found")
)

type UserRepo interface {
	Create(ctx context.Context, user models.User) (*models.User, error)
	Update(ctx context.Context, userId int, user models.User) (*models.User, error)
	Delete(ctx context.Context, userId int) (bool, error)
	GetByID(ctx context.Context, userId int) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetUsers(ctx context.Context, offset, page int) (users []*models.User, err error)
}

type UserRepoImpl struct {
	db     DBConnection
	logger log.SimpleLogger
}

func NewUserRepo(db DBConnection, logger log.SimpleLogger) UserRepo {
	return &UserRepoImpl{
		db:     db,
		logger: logger,
	}
}

func (uri *UserRepoImpl) Create(ctx context.Context, user models.User) (*models.User, error) {
	if result := uri.db.Insert(ctx, &user); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (uri *UserRepoImpl) Update(ctx context.Context, userId int, newUser models.User) (user *models.User, err error) {
	if user, err = uri.GetByID(ctx, userId); err != nil {
		return nil, err
	}

	user.Name = newUser.Name
	user.Address = newUser.Address
	user.Age = newUser.Age
	user.Email = newUser.Email
	if result := uri.db.Update(ctx, user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (uri *UserRepoImpl) Delete(ctx context.Context, userId int) (bool, error) {
	if _, err := uri.GetByID(ctx, userId); err != nil {
		return false, err
	}

	if result := uri.db.Delete(ctx, &models.User{}, userId); result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (uri *UserRepoImpl) GetByID(ctx context.Context, userId int) (*models.User, error) {
	user := &models.User{UserId: userId}
	if result := uri.db.First(ctx, user); result.Error != nil {
		if result.Error.Error() == RecordNotFoundErr.Error() {
			return nil, errors.New(fmt.Sprintf("User not found with id %d", userId))
		}

		return nil, result.Error
	}

	return user, nil
}
func (uri *UserRepoImpl) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	if result := uri.db.GetDB().WithContext(ctx).
		Where("email = ?", email).
		First(user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (uri *UserRepoImpl) GetUsers(ctx context.Context, offset, page int) (users []*models.User, err error) {
	if result := uri.db.FindAll(ctx, offset, page, "user_id", &users, "name", "age", "email", "address"); result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
