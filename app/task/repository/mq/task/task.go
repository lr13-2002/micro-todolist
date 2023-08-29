package task

import (
	"context"
	"encoding/json"
	"micro-todolist/app/task/repository/mq"
	"micro-todolist/app/task/service"
	"micro-todolist/consts"
	"micro-todolist/idl/pb"
)

type SyncTask struct {
}

func (s *SyncTask) RunTaskService(ctx context.Context) (err error) {
	rabbitMqQueue := consts.RabbitMqTaskQueue
	msgs, err := mq.ConnsumeMessage(ctx, rabbitMqQueue)
	if err != nil {
		return
	}
	var forever chan struct{}
	go func() {
		for d := range msgs {
			// 落库
			req := new(pb.TaskRequest)
			err := json.Unmarshal(d.Body, req)
			if err != nil {
				return
			}

			err = service.TaskMQ2DB(ctx, req)
			if err != nil {
				return
			}
		}
	}()
	<-forever
	return nil
}
