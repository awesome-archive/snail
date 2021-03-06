package main

import (
	"fmt"

	"github.com/LyricTian/snail/controllers"
	"github.com/LyricTian/snail/models"
	_ "github.com/LyricTian/snail/routers"
	"github.com/astaxie/beego"
)

func main() {
	// 初始化下载历史DB
	models.InitHistoryDB(beego.AppConfig.String("db::History"))
	defer models.HistoryDB.Close()

	// 初始化反馈建议DB
	models.InitSuggestDB(beego.AppConfig.String("db::Suggest"))
	defer models.SuggestDB.Close()

	// 设定日志
	beego.SetLogger("file", fmt.Sprintf(`{"filename":"%s"}`, beego.AppConfig.String("LogFile")))

	// 错误处理
	beego.ErrorController(&controllers.ErrorController{})

	beego.Run()
}
