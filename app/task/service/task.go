package service

import (
	"context"
	"encoding/json"
	"micro-todolist/app/task/repository/db/dao"
	"micro-todolist/app/task/repository/db/model"
	"micro-todolist/app/task/repository/mq"
	"micro-todolist/idl/pb"
	"micro-todolist/pkg/e"
	"sync"

	"github.com/gin-gonic/gin"
)

type TaskSrv struct {
}

var TaskSrvIns *TaskSrv
var TaskSrvOnce sync.Once

//懒汉式单例模式
func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

//create task 送到 mq mq ---> 落库
func (t *TaskSrv) CreateTask(ctx *gin.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = e.Success
	body, _ := json.Marshal(req)
	err = mq.SendMessage2MQ(body)
	if err != nil {
		resp.Code = e.Error
		return
	}
	return
}
func TaskMQ2DB(ctx context.Context, req *pb.TaskRequest) error {
	m := &model.Task{
		Uid:       uint(req.Uid),
		Title:     req.Title,
		Status:    int(req.Status),
		Content:   req.Content,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	return dao.NewTaskDao(ctx).CreateTask(m)
}
