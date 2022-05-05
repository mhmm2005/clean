package Logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
)

var Zapper *zap.Logger
var err error

func NewLogger() *zap.Logger {
	logPath := "/Users/momansouri/Desktop/logs/"
	config := zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	if runtime.GOOS != "windows" {
		err = os.MkdirAll(logPath, 0644)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		config.OutputPaths = []string{logPath + "clean.log"}
	} else {
		config.OutputPaths = []string{"clean.log"}
	}
	Zapper, err = config.Build()
	if err != nil {
		panic(err)
	}
	defer func(Zapper *zap.Logger) {
		err := Zapper.Sync()
		if err != nil {
			panic(err)
		}
	}(Zapper)

	return Zapper
}
