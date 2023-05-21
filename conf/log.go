package conf

import (
	"os"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() *zap.SugaredLogger {
	LogMode := zapcore.InfoLevel
	writeSyncer := getWriteSyncer()

	if viper.GetBool("app.debug") {
		writeSyncer = zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout))
		LogMode = zapcore.DebugLevel
	}

	core := zapcore.NewCore(getEncoder(), writeSyncer, LogMode)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	field := zap.Fields(zap.String("app_name", viper.GetString("app.name")))
	sugarLogger := zap.New(core, caller, development, field).Sugar()
	// 封装自己的日志打印方法
	logger := sugarLogger.Desugar().WithOptions(zap.AddCallerSkip(1)).Sugar()

	return logger
}

func getEncoder() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Local().Format(time.DateTime))
	}

	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriteSyncer() zapcore.WriteSyncer {
	lumberjackSync := &lumberjack.Logger{
		Filename:   viper.GetString("log.Dir") + time.Now().Format(time.DateOnly) + ".log",
		MaxSize:    viper.GetInt("log.MaxSize"),
		MaxBackups: viper.GetInt("log.MaxBackups"),
		MaxAge:     viper.GetInt("log.MaxAge"),
		Compress:   viper.GetBool("log.Compress"),
	}

	return zapcore.AddSync(lumberjackSync)
}
