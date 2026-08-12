package users

type IUserStore interface {
	Get(id UserId) (*UserDto, error)
	Register(user *CreateUserDto) (*UserDto, error)
	Login(user *LoginUserDto) (*LoginUserReturnDto, error)
	Update(id UserId, user *UpdateUserDto) (*UserDto, error)
	Delete(id UserId) error
}
