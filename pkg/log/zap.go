package log

import (
	"github.com/ljinf/template_project_v2/pkg/enum"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var _logger *zap.Logger

func NewLog(conf *viper.Viper) {

	// 日志地址 "out.log" 自定义
	lp := conf.GetString("log.log_file_name")
	// 日志级别 DEBUG,ERROR, INFO
	lv := conf.GetString("log.log_level")
	var level zapcore.Level
	//debug<info<warn<error<fatal<panic
	switch lv {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// 使用 lumberjack 实现 logger rotate
	hook := &lumberjack.Logger{
		Filename:   lp,                             // 日志文件路径
		MaxSize:    conf.GetInt("log.max_size"),    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: conf.GetInt("log.max_backups"), // 日志文件最多保存多少个备份
		MaxAge:     conf.GetInt("log.max_age"),     // 文件最多保存多少天
		Compress:   conf.GetBool("log.compress"),   // 是否压缩
		LocalTime:  true,
	}

	fileWriteSyncer := zapcore.AddSync(hook)

	var cores []zapcore.Core
	switch conf.GetString("app.env") {
	case enum.ModeTest, enum.ModeProd:
		// 测试环境和生产环境的日志输出到文件中
		cores = append(cores, zapcore.NewCore(encoder, fileWriteSyncer, level))
	case enum.ModeDev:
		// 开发环境同时向控制台和文件输出日志， Debug级别的日志也会被输出
		cores = append(
			cores,
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
			zapcore.NewCore(encoder, fileWriteSyncer, zapcore.DebugLevel),
		)

	}
	core := zapcore.NewTee(cores...)
	_logger = zap.New(core)
}
