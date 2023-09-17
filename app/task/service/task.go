package service

import (
	"context"
	"encoding/json"
	"log"
	"micro-todolist/app/task/repository/db/dao"
	"micro-todolist/app/task/repository/db/model"
	"micro-todolist/app/task/repository/db/redis"
	"micro-todolist/app/task/repository/mq"
	"micro-todolist/idl/pb"
	"micro-todolist/pkg/e"
	"sync"
	"time"

	Redis "github.com/go-redis/redis/v8"
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
func (*TaskSrv) CreateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
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

func (*TaskSrv) GetTasksList(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskListResponse) (err error) {
	resp.Code = e.Success
	if req.Limit == 0 {
		req.Limit = 10
	}
	r, err := redis.GetListTask(req.Uid, ctx)
	if err != nil {
		if err == Redis.Nil {
			log.Printf("走的db")
			r, err = dao.NewTaskDao(ctx).ListTaskByUserId(req.Uid)
			if err != nil {
				resp.Code = e.Error
				return
			}

			err = redis.SetListTask(req.Uid, r, ctx)
			if err != nil {
				resp.Code = e.Error
				return
			}

		} else {
			if err != nil {
				resp.Code = e.Error
				return
			}
		}
	} else {
		log.Printf("走的缓存")
	}

	var taskRes []*pb.TaskModel
	for ind, item := range r {
		if ind >= int(req.Start) && ind < int(req.Limit) {
			taskRes = append(taskRes, BuildTask(item))
		}
	}
	resp.TaskList = taskRes
	resp.Count = uint32(len(taskRes))

	return
}

func (*TaskSrv) GetTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = e.Success

	r, err := redis.GetTask(req.Uid, req.Id, ctx)
	if err != nil {
		if err == Redis.Nil {
			log.Printf("走的db")
			r, err = dao.NewTaskDao(ctx).GetTaskByTaskAndUserId(req.Id, req.Uid)
			if r.ID == 0 || err != nil {
				resp.Code = e.Error
				return
			}

			err = redis.SetTask(req.Uid, req.Id, r, ctx)
			if err != nil {
				resp.Code = e.Error
				return
			}

		} else {
			if err != nil {
				resp.Code = e.Error
				return
			}
		}
	} else {
		log.Printf("走的缓存")
	}
	resp.TaskDetail = BuildTask(r)
	return
}

func (*TaskSrv) UpdateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = e.Success

	result, err := redis.DelTask(req.Uid, req.Id, ctx)
	if err != nil {
		resp.Code = e.Error
		return
	}

	log.Println("删除了", result)

	result, err = redis.DelListTask(req.Uid, ctx)
	if err != nil {
		resp.Code = e.Error
		return
	}

	log.Println("删除了", result)

	err = dao.NewTaskDao(ctx).UpdateTask(req)
	if err != nil {
		resp.Code = e.Error
		return
	}

	time.Sleep(2 * time.Millisecond)
	result, err = redis.DelTask(req.Uid, req.Id, ctx)
	if err != nil {
		resp.Code = e.Error
		return
	}

	log.Println("删除了", result)

	result, err = redis.DelListTask(req.Uid, ctx)
	if err != nil {
		resp.Code = e.Error
		return
	}

	log.Println("删除了", result)

	return
}

func (*TaskSrv) DeleteTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = e.Success

	result, err := redis.DelTask(req.Uid, req.Id, ctx)
	if err != nil {
		resp.Code = e.Error
		return
	}

	log.Println("删除了", result)

	result, err = redis.DelListTask(req.Uid, ctx)
	if err != nil {
		resp.Code = e.Error
		return
	}

	log.Println("删除了", result)

	err = dao.NewTaskDao(ctx).DeleteTaskByTaskAndUserId(req.Id, req.Uid)
	if err != nil {
		resp.Code = e.Error
		return
	}

	time.Sleep(2 * time.Millisecond)
	result, err = redis.DelTask(req.Uid, req.Id, ctx)
	if err != nil {
		resp.Code = e.Error
		return
	}

	log.Println("删除了", result)

	result, err = redis.DelListTask(req.Uid, ctx)
	if err != nil {
		resp.Code = e.Error
		return
	}

	log.Println("删除了", result)

	return
}

func BuildTask(item *model.Task) *pb.TaskModel {
	return &pb.TaskModel{
		Id:         uint64(item.ID),
		Uid:        uint64(item.Uid),
		Title:      item.Title,
		Content:    item.Content,
		StartTime:  item.StartTime,
		EndTime:    item.EndTime,
		Status:     int64(item.Status),
		CreateTime: item.CreatedAt.Unix(),
		UpdateTime: item.UpdatedAt.Unix(),
	}
}
