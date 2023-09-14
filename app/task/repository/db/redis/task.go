package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"micro-todolist/app/task/repository/db/model"
	"strconv"
	"time"
)

func SetListTask(userId uint64, value []*model.Task, ctx context.Context) (err error) {
	ctx_, cancel_ := context.WithTimeout(ctx, 2*time.Second)
	defer cancel_()

	value_, err := json.Marshal(value)
	if err != nil {
		return
	}

	err = redisClient.Set(ctx_, strconv.Itoa(int(userId)), value_, 100*time.Second).Err()
	if err != nil {
		return
	}

	return
}

func SetTask(userId uint64, taskId uint64, value *model.Task, ctx context.Context) (err error) {
	ctx_, cancel_ := context.WithTimeout(ctx, 2*time.Second)
	defer cancel_()

	value_, err := json.Marshal(value)
	if err != nil {
		return
	}

	err = redisClient.Set(ctx_, fmt.Sprintf("%s.%s", strconv.Itoa(int(userId)), strconv.Itoa(int(taskId))), value_, 100*time.Second).Err()
	if err != nil {
		return
	}

	return
}

func GetListTask(userId uint64, ctx context.Context) (r []*model.Task, err error) {
	ctx_, cancel_ := context.WithTimeout(ctx, 2*time.Second)
	defer cancel_()

	result, err := redisClient.Get(ctx_, strconv.Itoa(int(userId))).Bytes()
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &r)
	return
}

func DelListTask(userId uint64, ctx context.Context) (result int64, err error) {
	ctx_, cancel_ := context.WithTimeout(ctx, 2*time.Second)
	defer cancel_()

	result, err = redisClient.Del(ctx_, strconv.Itoa(int(userId))).Result()
	return
}

func GetTask(userId uint64, taskId uint64, ctx context.Context) (r *model.Task, err error) {
	ctx_, cancel_ := context.WithTimeout(ctx, 2*time.Second)
	defer cancel_()

	result, err := redisClient.Get(ctx_, fmt.Sprintf("%s.%s", strconv.Itoa(int(userId)), strconv.Itoa(int(taskId)))).Bytes()
	if err != nil {
		return
	}

	err = json.Unmarshal(result, &r)
	if err != nil {
		return
	}

	return
}

func DelTask(userId uint64, taskId uint64, ctx context.Context) (result int64, err error) {
	ctx_, cancel_ := context.WithTimeout(ctx, 2*time.Second)
	defer cancel_()

	result, err = redisClient.Del(ctx_, fmt.Sprintf("%s.%s", strconv.Itoa(int(userId)), strconv.Itoa(int(taskId)))).Result()
	return
}
