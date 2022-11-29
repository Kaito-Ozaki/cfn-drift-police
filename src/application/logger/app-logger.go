package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"cfn-drift-police/src/application/consts"
)

/**
AppLogger はアプリケーション共通のロガーです。
*/
type AppLogger struct {
	Logger *zap.Logger
}

/**
NewAppLoger はAppLoggerを生成するための、初期化関数です。
*/
func NewAppLoger() AppLogger {
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)
	myConfig := zap.Config{
		Level:    level,
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "Time",
			LevelKey:       "Level",
			NameKey:        "Name",
			CallerKey:      "Caller",
			MessageKey:     "Msg",
			StacktraceKey:  "St",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, _ := myConfig.Build()

	return AppLogger{
		Logger: logger,
	}
}

/**
Infoは、ログレベルINFOのログを出力する関数です。
*/
func (n *AppLogger) Info(l consts.Log) {
	n.Logger.Info(l.Message, zap.String(consts.LOG_ID, l.Id))
}

/**
InfoWithParamsは、ログレベルINFOのログを出力する関数です。
引数を伴うログを出力したい際に使用します。
*/
func (n *AppLogger) InfoWithParams(l consts.Log, args ...interface{}) {
	n.Logger.Info(fmt.Sprintf(l.Message, args), zap.String(consts.LOG_ID, l.Id))
}

/**
Errorは、ログレベルERRORのログを出力する関数です。
*/
func (n *AppLogger) Error(l consts.Log, e error) {
	n.Logger.Error(fmt.Sprintf(l.Message, e.Error()), zap.String(consts.LOG_ID, l.Id))
}

/**
Errorは、ログレベルERRORのログを出力する関数です。
引数を伴うログを出力したい際に使用します。
*/
func (n *AppLogger) ErrorWithParams(l consts.Log, e error, args ...interface{}) {
	n.Logger.Error(fmt.Sprintf(l.Message, args, e.Error()), zap.String(consts.LOG_ID, l.Id))
}

/**
Fatalは、ログレベルFATALのログを出力し、該当処理が行われているプロセスを終了する関数です。
*/
func (n *AppLogger) Fatal(l consts.Log, e error) {
	n.Logger.Fatal(fmt.Sprintf(l.Message, e.Error()), zap.String(consts.LOG_ID, l.Id))
}

/**
FatalWithParamsは、ログレベルFATALのログを出力し、該当処理が行われているプロセスを終了する関数です。
引数を伴うログを出力したい際に使用します。
*/
func (n *AppLogger) FatalWithParams(l consts.Log, e error, args ...interface{}) {
	n.Logger.Fatal(fmt.Sprintf(l.Message, args, e.Error()), zap.String(consts.LOG_ID, l.Id))
}
