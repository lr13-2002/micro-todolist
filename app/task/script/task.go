package script

import (
	"context"
	"micro-todolist/app/task/repository/mq/task"
)

func TaskCreateSync(ctx context.Context) {
	tSync := new(task.SyncTask)
	err := tSync.RunTaskService(ctx)
	if err != nil {
		return 
	}
}
