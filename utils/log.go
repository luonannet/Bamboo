package utils

import "github.com/astaxie/beego/logs"

//Log 日志
var logger *logs.BeeLogger

func init() {
	logger = logs.NewLogger(logs.LevelDebug)
	logger.SetLogger("multifile", `{"filename":"log/test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	// log.Debug("debug")
	// log.Informational("info")
	// log.Notice("notice")
	// log.Warning("warning")
	// log.Error("error")
	// log.Alert("alert")
	// log.Critical("critical")
	// log.Emergency("emergency")

}

//Debug Debug
func Debug(info ...interface{}) {
	logger.Debug("debug: %v", info)
}

//Error Error
func Error(info ...interface{}) {
	logger.Error("error: %v", info)
}
