package dao

import "micro-todolist/app/user/repositery/db/model"

func migration() {
	_db.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&model.User{})
}
