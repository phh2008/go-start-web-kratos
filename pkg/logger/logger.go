package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"helloword/internal/conf"
	"io"
	"os"
	"strings"
	"time"
)

var log *zap.Logger = zap.L()
var slog *zap.SugaredLogger = log.Sugar()
var wrapLog = copyLog(zap.L()).WithOptions(zap.AddCallerSkip(1))
var wrapSugarLog = wrapLog.Sugar()

var levelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

// NewLogger 创建 logger
func NewLogger(config *conf.Log) *zap.Logger {
	log = newZapLog(config)
	slog = log.Sugar()
	// 替换 zap 全局 log
	zap.ReplaceGlobals(log)
	// 包装 log
	wrapLog = copyLog(log).WithOptions(zap.AddCallerSkip(1))
	wrapSugarLog = wrapLog.Sugar()
	return log
}

func copyLog(log *zap.Logger) *zap.Logger {
	l := *log
	return &l
}

func L() *zap.Logger {
	l := log
	return l
}

func S() *zap.SugaredLogger {
	s := slog
	return s
}

// Debug as zap.L().Debug
func Debug(msg string, fields ...zap.Field) {
	wrapLog.Debug(msg, fields...)
}

// Info as zap.L().Info
func Info(msg string, fields ...zap.Field) {
	wrapLog.Info(msg, fields...)
}

// Warn as zap.L().Warn
func Warn(msg string, fields ...zap.Field) {
	wrapLog.Warn(msg, fields...)
}

// Error as zap.L().Error
func Error(msg string, fields ...zap.Field) {
	wrapLog.Error(msg, fields...)
}

// Debugf as zap.S().Debugf
func Debugf(template string, args ...interface{}) {
	wrapSugarLog.Debugf(template, args...)
}

// Infof zap.S().Infof
func Infof(template string, args ...interface{}) {
	wrapSugarLog.Infof(template, args...)
}

// Warnf zap.S().Warnf
func Warnf(template string, args ...interface{}) {
	wrapSugarLog.Warnf(template, args...)
}

// Errorf as zap.S().Errorf
func Errorf(template string, args ...interface{}) {
	wrapSugarLog.Errorf(template, args...)
}

// Debugs uses fmt.Sprint to construct and log a message.
func Debugs(args ...interface{}) {
	wrapSugarLog.Debug(args...)
}

// Infos uses fmt.Sprint to construct and log a message.
func Infos(args ...interface{}) {
	wrapSugarLog.Info(args...)
}

// Warns uses fmt.Sprint to construct and log a message.
func Warns(args ...interface{}) {
	wrapSugarLog.Warn(args...)
}

// Errors uses fmt.Sprint to construct and log a message.
func Errors(args ...interface{}) {
	wrapSugarLog.Error(args...)
}

// getWriter
func getWriter(config *conf.Log) io.Writer {
	return &lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    int(config.MaxSize), // megabytes
		MaxBackups: int(config.MaxBackups),
		MaxAge:     int(config.MaxAge), //days
		LocalTime:  config.LocalTime,
		Compress:   config.Compress, // disabled by default
	}
}

func newZapLog(config *conf.Log) *zap.Logger {
	// 设置日志格式
	//encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	// 记录什么级别的日志
	level := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= levelMap[strings.ToLower(config.Level)]
	})

	// 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现
	writer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(getWriter(config)))
	// 如果info、debug、error分文件记录，就创建多个 writer
	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writer, level), // 可添加多个
	)
	// 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
	return zap.New(core, zap.AddCaller())
}
