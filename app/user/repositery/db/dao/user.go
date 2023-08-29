package dao

import (
	"context"
	"micro-todolist/app/user/repositery/db/model"

	"gorm.io/gorm"
)

//定义得是对数据库得 user model 得 crud 操作
type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) FindUserByUserName(userName string) (r *model.User, err error) {
	err = dao.Model(&model.User{}).Where("user_name = ? ", userName).Find(&r).Error

	return
}

func (dao *UserDao) CreateUser(in *model.User) (err error) {
	return dao.Model(&model.User{}).Create(&in).Error
}
