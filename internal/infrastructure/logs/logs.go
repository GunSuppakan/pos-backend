package logs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/patcharp/golib/requests"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
	pushToLoki("info", message)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), fields...)
		pushToLoki("error", v.Error())
	case string:
		log.Error(v, fields...)
		pushToLoki("error", v)
	}
}

type modelLokiApi struct {
	Streams []bodyStreams `json:"streams"`
}

type bodyStreams struct {
	Labels  string    `json:"labels"`
	Entries []logData `json:"entries"`
}

type logData struct {
	Ts   string `json:"ts"`
	Line string `json:"line"`
}

func pushToLoki(levelLog string, message string) {
	if os.Getenv("ENV") != "pro" && os.Getenv("ENV") != "prd" && os.Getenv("ENV") != "uat" {
		return
	}
	_, file, no, ok := runtime.Caller(2)
	if ok {
		now := time.Now()
		timeStamp := now.Format(time.RFC3339)
		lineLog := fmt.Sprintf("level: %v caller:%v:%v msg: %v", levelLog, file, no, message)
		go func() {
			headers := map[string]string{
				"Accept":       "application/json",
				"Content-Type": "application/json",
			}

			myLogData := logData{
				Ts:   timeStamp,
				Line: lineLog,
			}
			listLogs := []logData{myLogData}

			myBodySteam := bodyStreams{
				Labels:  fmt.Sprintf("{service=\"%v\"}", viper.GetString("loki.service")),
				Entries: listLogs,
			}

			data := modelLokiApi{
				Streams: []bodyStreams{myBodySteam},
			}

			body, _ := json.Marshal(&data)

			res, err := requests.Post(fmt.Sprintf("%v/api/prom/push", viper.GetString("loki.url")), headers, bytes.NewBuffer(body), 10)
			if err != nil {
				fmt.Println(err)
			}
			_ = res

		}()
	}
}
