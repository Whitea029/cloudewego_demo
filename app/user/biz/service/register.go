package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Whitea029/whmall/app/user/biz/dal/mysql"
	"github.com/Whitea029/whmall/app/user/biz/model"
	user "github.com/Whitea029/whmall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// TODO verify email format
	if req.Email == "" || req.Password == "" || req.PasswordConfirm == "" {
		return nil, errors.New("email, password, password confirm can not be empty")
	}
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("password and password confirm not match")
	}
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwordHashed),
	}
	err = model.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}
	fmt.Println("newUser.ID", newUser.ID)
	return &user.RegisterResp{
		UserId: int32(newUser.ID),
	}, nil
}
