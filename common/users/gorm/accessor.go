package gormusers

import (
	"github.com/criptycpizza7/habit-tracker/common/db"
	gormdb "github.com/criptycpizza7/habit-tracker/common/db/gorm_db"
	"github.com/criptycpizza7/habit-tracker/common/users"
	usermodels "github.com/criptycpizza7/habit-tracker/common/users/gorm/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserAccessor struct {
	db *gormdb.DB
}

func NewUserAccessor(db *gormdb.DB) *UserAccessor {
	return &UserAccessor{db}
}

func fromGormToDto(user *usermodels.User) *users.UserDto {
	return &users.UserDto{
		ID:        users.UserId(user.ID),
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (a *UserAccessor) Get(id users.UserId) (*users.UserDto, error) {
	user := &usermodels.User{}
	result := a.db.First(user, "id = ?::uuid", id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, db.ErrUserNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return fromGormToDto(user), nil
}

func (a *UserAccessor) Register(user *users.CreateUserDto) (*users.UserDto, error) {
	user_model := &usermodels.User{
		Username: user.Username,
		Password: user.Password,
	}
	result := a.db.Create(user_model).Clauses(clause.Returning{})
	if result.Error == gorm.ErrDuplicatedKey {
		return nil, db.ErrUserAlreadyExists
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return fromGormToDto(user_model), nil
}

func (a *UserAccessor) Login(user *users.LoginUserDto) (*users.LoginUserReturnDto, error) {
	user_model := &usermodels.User{}
	result := a.db.First(user_model, "username = ?", user.Username)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, db.ErrUserNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &users.LoginUserReturnDto{
		ID:        users.UserId(user_model.ID),
		Username:  user_model.Username,
		CreatedAt: user_model.CreatedAt,
		UpdatedAt: user_model.UpdatedAt,
		Password:  user_model.Password,
	}, nil
}

func (a *UserAccessor) Update(id users.UserId, user *users.UpdateUserDto) (*users.UserDto, error) {
	user_model := &usermodels.User{}
	result := a.db.Model(user_model).First(&usermodels.User{ID: uuid.UUID(id)}).Clauses(clause.Returning{}).Updates(*user)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, db.ErrUserNotFound
	}
	if result.Error == gorm.ErrDuplicatedKey {
		return nil, db.ErrUserAlreadyExists
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return fromGormToDto(user_model), nil
}

func (a *UserAccessor) Delete(id users.UserId) error {
	result := a.db.Delete(&usermodels.User{}, "id = ?", id)
	if result.Error == gorm.ErrRecordNotFound {
		return db.ErrUserNotFound
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
