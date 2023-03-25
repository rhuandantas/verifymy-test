package repo

import (
	"context"
	"github.com/rhuandantas/verifymy-test/internal/log"
	"github.com/rhuandantas/verifymy-test/internal/models"
)

//go:generate mockgen -source=$GOFILE -package=mock_repo -destination=../../test/mock/repo/$GOFILE

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
	if result := uri.db.GetDB().WithContext(ctx).Create(&user); result.Error != nil {
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
	if result := uri.db.GetDB().WithContext(ctx).Save(user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (uri *UserRepoImpl) Delete(ctx context.Context, userId int) (bool, error) {
	if _, err := uri.GetByID(ctx, userId); err != nil {
		return false, err
	}

	if result := uri.db.GetDB().WithContext(ctx).Delete(&models.User{}, userId); result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (uri *UserRepoImpl) GetByID(ctx context.Context, userId int) (*models.User, error) {
	user := &models.User{UserId: userId}
	if result := uri.db.GetDB().WithContext(ctx).First(user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
func (uri *UserRepoImpl) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	if result := uri.db.GetDB().WithContext(ctx).
		Where("email = ?", email).
		Find(user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (uri *UserRepoImpl) GetUsers(ctx context.Context, offset, page int) (users []*models.User, err error) {
	if result := uri.db.GetDB().WithContext(ctx).
		Offset(offset).
		Limit(page).
		Select("user_id", "name", "age", "email", "address").
		Find(&users); result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
