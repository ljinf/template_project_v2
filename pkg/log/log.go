package log

import (
	"context"
	"fmt"
	"github.com/ljinf/template_project_v2/pkg/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"runtime"
)

func Debug(ctx context.Context, msg string, kv ...interface{}) {
	log(zapcore.DebugLevel, msg, append(kv, traceInfo(ctx)...)...)
}

func Info(ctx context.Context, msg string, kv ...interface{}) {
	log(zapcore.InfoLevel, msg, append(kv, traceInfo(ctx)...)...)
}

func Warn(ctx context.Context, msg string, kv ...interface{}) {
	log(zapcore.WarnLevel, msg, append(kv, traceInfo(ctx)...)...)
}

func Error(ctx context.Context, msg string, kv ...interface{}) {
	log(zapcore.ErrorLevel, msg, append(kv, traceInfo(ctx)...)...)
}

func Fatal(ctx context.Context, msg string, kv ...interface{}) {
	log(zapcore.FatalLevel, msg, append(kv, traceInfo(ctx)...)...)
	os.Exit(1)
}

func traceInfo(ctx context.Context) []interface{} {
	// 日志行信息中增加追踪参数
	list := make([]interface{}, 0, 6)
	traceId, spanId, pSpanId := util.GetTraceInfoFromCtx(ctx)
	list = append(list, "traceid", traceId, "spanid", spanId, "pspanid", pSpanId)
	return list
}

// kv 应该是成对的数据, 类似: name,张三,age,10,...
func log(lvl zapcore.Level, msg string, kv ...interface{}) {

	if _logger == nil {
		return
	}

	//调用zap.check判断这个日志级别能否写入
	if ce := _logger.Check(lvl, msg); ce != nil {

		// 保证要打印的日志信息成对出现
		if len(kv)%2 != 0 {
			kv = append(kv, "unknown")
		}

		// 增加日志调用者信息, 方便查日志时定位程序位置
		funcName, file, line := getLoggerCallerInfo()
		kv = append(kv, "func", funcName, "file", file, "line", line)

		fields := make([]zap.Field, 0, len(kv)/2)
		for i := 0; i < len(kv); i += 2 {
			k := fmt.Sprintf("%v", kv[i])
			fields = append(fields, zap.Any(k, kv[i+1]))
		}

		ce.Write(fields...)
	}
}

// getLoggerCallerInfo 日志调用者信息 -- 方法名, 文件名, 行号
func getLoggerCallerInfo() (funcName, file string, line int) {

	pc, file, line, ok := runtime.Caller(3) // 回溯拿调用日志方法的业务函数的信息
	if !ok {
		return
	}
	file = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	return
}
