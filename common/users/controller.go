package users

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userStore IUserStore
}

func NewController(userStore IUserStore) *UserController {
	return &UserController{userStore}
}

func (u *UserController) Get(id UserId) (*UserDto, error) {
	return u.userStore.Get(id)
}

func (u *UserController) hashPassword(password string) (string, error) {
	new_password, err := bcrypt.GenerateFromPassword([]byte(password), HASH_COST)
	if err != nil {
		return "", err
	}
	return string(new_password), nil
}

func (u *UserController) Register(user *CreateUserDto) (*UserDto, error) {
	var err error
	user.Password, err = u.hashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	return u.userStore.Register(user)
}

func (u *UserController) Login(user *LoginUserDto) (*string, error) {
	login_user, err := u.userStore.Login(user)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(login_user.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}

	claims := JWTClaims{
		login_user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWT_TOKEN_TTL)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token_string, err := token.SignedString(JWT_SECRET)

	return &token_string, err
}

func (u *UserController) Update(id UserId, user *UpdateUserDto) (*UserDto, error) {
	return u.userStore.Update(id, user)
}

func (u *UserController) Delete(id UserId) error {
	return u.userStore.Delete(id)
}
