package zap

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "google.golang.org/grpc/grpclog"
    "gopkg.in/natefinch/lumberjack.v2"
    "github.com/go-kit/kit/log"
)

//LoggerConfig config of logger
type LoggerConfig struct {
    Level      string `yaml:"level"`       //debug  info  warn  error
    Encoding   string `yaml:"encoding"`    //json or console
    CallFull   bool   `yaml:"call_full"`   //whether full call path or short path, default is short
    Filename   string `yaml:"file_name"`   //log file name
    MaxSize    int    `yaml:"max_size"`    //max size of log.(MB)
    MaxAge     int    `yaml:"max_age"`     //time to keep, (day)
    MaxBackups int    `yaml:"max_backups"` //max file numbers
    LocalTime  bool   `yaml:"local_time"`  //(default UTC)
    Compress   bool   `yaml:"compress"`    //default false
}

func convertLogLevel(levelStr string) (level zapcore.Level) {
    switch levelStr {
    case "debug":
        level = zap.DebugLevel
    case "info":
        level = zap.InfoLevel
    case "warn":
        level = zap.WarnLevel
    case "error":
        level = zap.ErrorLevel
    }
    return
}

//NewDefaultLoggerConfig create a default config
func NewDefaultLoggerConfig() *LoggerConfig {
    return &LoggerConfig{
        Level:      "debug",
        MaxSize:    1,
        MaxAge:     1,
        MaxBackups: 10,
    }
}

var atom zap.AtomicLevel

// 日志时间格式
// func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
// 	enc.AppendString(t.Format("2006-01-02 15:04:05"))
// }
//NewLogger create logger by config
func (lconf *LoggerConfig) NewLogger() *ZapLogger {
    if lconf.Filename == "" {
        config := zap.NewProductionConfig()
        config.DisableCaller = true
        logger, _ := config.Build()
        //zap.NewProduction(zap.AddCallerSkip(2))
        return NewZapLogger(logger)
    }

    enCfg := zap.NewProductionEncoderConfig()
    if lconf.CallFull {
        enCfg.EncodeCaller = zapcore.FullCallerEncoder
    }
    encoder := zapcore.NewJSONEncoder(enCfg)
    if lconf.Encoding == "console" {
        zapcore.NewConsoleEncoder(enCfg)
    }

    //zapWriter := zapcore.
    zapWriter := zapcore.AddSync(&lumberjack.Logger{
        Filename:   lconf.Filename,
        MaxSize:    lconf.MaxSize,
        MaxAge:     lconf.MaxAge,
        MaxBackups: lconf.MaxBackups,
        LocalTime:  lconf.LocalTime,
    })

    atom = zap.NewAtomicLevel()
    atom.SetLevel(convertLogLevel(lconf.Level))
    newCore := zapcore.NewCore(encoder, zapWriter, atom)
    opts := []zap.Option{zap.ErrorOutput(zapWriter)}
    //opts = append(opts, zap.AddCaller(), zap.AddCallerSkip(2))
    logger := zap.New(newCore, opts...)
    return NewZapLogger(logger)
}

func AutoSetLogLevel(level  string){
    atom.SetLevel(convertLogLevel(level))
}

type DefaultLogger struct{}

// NewNopLogger returns a logger that doesn't do anything.
func NewDefaultLogger() log.Logger { return DefaultLogger{} }

func (DefaultLogger) Log(keyvals ...interface{}) error {
    grpclog.Error(keyvals)
    return nil
}


