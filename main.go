package main

import (
	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	log "github.com/sirupsen/logrus"
)

func main() {
	hook := graylog.NewGraylogHook("192.168.0.118:12201", map[string]interface{}{"this": "is logged every time"})
	log.AddHook(hook)
	log.Info("some logging message")
}
