package dao

import (
	"context"
	"micro-todolist/app/task/repository/db/model"

	"gorm.io/gorm"
)

//定义得是对数据库得 task model 得 crud 操作
type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{NewDBClient(ctx)}
}

func (dao *TaskDao) CreateTask(data *model.Task) error {
	return dao.Model(&model.Task{}).Create(&data).Error
}
