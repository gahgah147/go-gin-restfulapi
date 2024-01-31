package main

import (
	"fmt"
	//"message/model"
	"message/model"
	"message/router"
	"message/sql"
)

func main() {
	//連線資料庫
	if err := sql.InitMySql(); err != nil {
		panic(err)
	}

	//連結模型
	sql.Connect.AutoMigrate(&model.Message{})
	sql.Connect.Table("message")

	//註冊路由
	r := router.SetRouter()

	//啟動埠為8081的專案
	fmt.Println("開啟:8082...")
	r.Run(":8082")
}
