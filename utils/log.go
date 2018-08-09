package utils

import "github.com/astaxie/beego/logs"

//Log 日志
var log *logs.BeeLogger

func init() {
	log = logs.NewLogger(logs.LevelDebug)
	log.SetLogger("multifile", `{"filename":"log/test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	// log.Debug("debug")
	// log.Informational("info")
	// log.Notice("notice")
	// log.Warning("warning")
	// log.Error("error")
	// log.Alert("alert")
	// log.Critical("critical")
	// log.Emergency("emergency")

}
func Debug(info string) {
	log.Debug("")
}
func Error(info string) {
	log.Error("")
}
