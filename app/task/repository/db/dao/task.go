package dao

import (
	"context"
	"log"
	"micro-todolist/app/task/repository/db/model"
	"micro-todolist/idl/pb"

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

func (dao *TaskDao) ListTaskByUserId(userId uint64) (r []*model.Task, err error) {
	err = dao.Model(&model.Task{}).Where("uid = ?", userId).Find(&r).Error
	if err != nil {
		return
	}

	return
}

func (dao *TaskDao) GetTaskByTaskAndUserId(tId, uId uint64) (r *model.Task, err error) {
	err = dao.Model(&model.Task{}).Where("id = ? AND uid = ?", tId, uId).Find(&r).Error
	return
}

func (dao *TaskDao) DeleteTaskByTaskAndUserId(tId, uId uint64) (err error) {
	err = dao.Model(&model.Task{}).Where("id = ? AND uid = ?", tId, uId).Delete(&model.Task{}).Error
	return
}

func (dao *TaskDao) UpdateTask(req *pb.TaskRequest) (err error) {
	var r *model.Task
	err = dao.Model(&model.Task{}).Where("id = ? AND uid = ?", req.Id, req.Uid).Find(&r).Error
	if err != nil {
		return
	}
	r.Title = req.Title
	r.Status = int(req.Status)
	r.Content = req.Content
	log.Println("更新操作 ", r.Title, r.Status, r.Content, r.ID)
	return dao.Save(&r).Error
}
