package service

import (
	"context"
	"errors"
	"micro-todolist/app/user/repositery/db/dao"
	"micro-todolist/app/user/repositery/db/model"
	"micro-todolist/idl/pb"
	"micro-todolist/pkg/e"
	"sync"

	"gorm.io/gorm"
)

type UserSrv struct {
}

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

//懒汉式单例模式
func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (u *UserSrv) UserLogin(ctx context.Context, req *pb.UserRequest, resp *pb.UserResponse) (err error) {
	resp.Code = e.Success
	//查看有没有这个人
	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		return
	}
	if user.ID == 0 {
		err = errors.New("用户不存在")
		return
	}

	if !user.CheckPassword(req.Password) {
		err = errors.New("用户密码错误")
		resp.Code = e.Error
		return
	}

	resp.UserDetail = BuildUser(user)
	return
}

func (u *UserSrv) UserRegister(ctx context.Context, req *pb.UserRequest, resp *pb.UserResponse) (err error) {
	resp.Code = e.Success

	if req.Password != req.PasswordConfirm {
		err = errors.New("两次密码不一致")
		resp.Code = e.Error
		return
	}

	if req.UserName == "" {
		err = errors.New("用户名为空")
		resp.Code = e.Error
		return
	}
	_, err = dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		if err == gorm.ErrRecordNotFound {

		} else {
			resp.Code = e.Error
			return
		}
	}

	user := &model.User{
		UserName: req.UserName,
	}

	if user.ID > 0 {
		err = errors.New("用户名已存在")
		resp.Code = e.Error
		return
	}

	user = &model.User{
		UserName: req.UserName,
	}

	// 加密密码
	if err = user.SetPassword(req.Password); err != nil {
		resp.Code = e.Error
		return
	}

	if err = dao.NewUserDao(ctx).CreateUser(user); err != nil {
		resp.Code = e.Error
		return
	}

	resp.UserDetail = BuildUser(user)
	return
}

func BuildUser(item *model.User) *pb.UserModel {
	return &pb.UserModel{
		Id:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.CreatedAt.Unix(),
	}
}
