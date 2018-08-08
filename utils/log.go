package utils

import "github.com/astaxie/beego/logs"

//Log 日志
var Log *logs.BeeLogger

func init() {
	Log = logs.NewLogger(logs.LevelDebug)
	Log.SetLogger("multifile", `{"filename":"test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	Log.Debug("debug")
	Log.Informational("info")
	Log.Notice("notice")
	Log.Warning("warning")
	Log.Error("error")
	Log.Alert("alert")
	Log.Critical("critical")
	Log.Emergency("emergency")

}
