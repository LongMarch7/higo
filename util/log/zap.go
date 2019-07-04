package zap

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "google.golang.org/grpc/grpclog"
    "gopkg.in/natefinch/lumberjack.v2"
    "github.com/go-kit/kit/log"
)

//LoggerConfig config of logger
const TYPE_FILE = "file"
const TYPE_Console = "console"
type Config struct {
    Type       string `json:"type"`         //1 - out file  2 out consol
    Level      string `json:"level"`       //debug  info  warn  error
    Encoding   string `json:"encoding"`    //json or console
    CallFull   bool   `json:"call_full"`   //whether full call path or short path, default is short
    Filename   string `json:"file_name"`   //log file name
    MaxSize    int    `json:"max_size"`    //max size of log.(MB)
    MaxAge     int    `json:"max_age"`     //time to keep, (day)
    MaxBackups int    `json:"max_backups"` //max file numbers
    LocalTime  bool   `json:"local_time"`  //(default UTC)
    Compress   bool   `json:"compress"`    //default false
}

func convertLogLevel(levelStr string) (level zapcore.Level) {
    switch levelStr {
    case "debug":
        fallthrough
    case "DEBUG":
        level = zap.DebugLevel
    case "info":
        fallthrough
    case "INFO":
        level = zap.InfoLevel
    case "warn":
        fallthrough
    case "WARN":
        level = zap.WarnLevel
    case "error":
        fallthrough
    case "ERROR":
        level = zap.ErrorLevel
    default:
        level = zap.ErrorLevel
    }
    return
}

//NewDefaultLoggerConfig create a default config
func NewDefaultLoggerConfig() *Config {
    return &Config{
        Level:      "debug",
        MaxSize:    1,
        MaxAge:     1,
        MaxBackups: 10,
        LocalTime: true,
    }
}

var atom zap.AtomicLevel

// 日志时间格式
// func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
// 	enc.AppendString(t.Format("2006-01-02 15:04:05"))
// }
//NewLogger create logger by config
func (lconf *Config) NewLogger() *ZapLogger {
    if lconf.Filename == "" {
        config := zap.NewProductionConfig()
        config.DisableCaller = true
        config.Level =  zap.NewAtomicLevelAt(convertLogLevel(lconf.Level))
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

func SetLogLevel(level  string){
    atom.SetLevel(convertLogLevel(level))
}

type DefaultLogger struct{}

// NewNopLogger returns a logger that doesn't do anything.
func NewDefaultLogger() log.Logger { return DefaultLogger{} }

func (DefaultLogger) Log(keyvals ...interface{}) error {
    grpclog.Error(keyvals)
    return nil
}


