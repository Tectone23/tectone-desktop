package main

import (
	"context"
	"encoding/json"
	"os"
	"runtime"
	"time"
)

const (
	LogInfo  LogLevel = "INFO"
	LogDebug LogLevel = "DEBUG"
	LogError LogLevel = "ERROR"
	LogFatal LogLevel = "FATAL"
)

type LogLevel string

type Log struct {
	Message   string   `json:"message"`
	Level     LogLevel `json:"level"`
	Timestamp int64    `json:"timestamp"`
	Os        string   `json:"os"`
}

type Logger struct {
	ctx context.Context
	f   *os.File
}

func (l *Logger) Info(msg string) {
	log := &Log{
		Message:   msg,
		Level:     LogInfo,
		Timestamp: time.Now().UnixMilli(),
		Os:        runtime.GOOS,
	}

	b, err := json.Marshal(log)
	if err != nil {
		// TODO: handle error properly
	}

	if err := l.write(b); err != nil {
		// TODO: handle error properly
	}
}

func (l *Logger) Debug(msg string) {
	log := &Log{
		Message:   msg,
		Level:     LogDebug,
		Timestamp: time.Now().UnixMilli(),
		Os:        runtime.GOOS,
	}

	b, err := json.Marshal(log)
	if err != nil {
		// TODO: handle error properly
	}

	if err := l.write(b); err != nil {
		// TODO: handle error properly
	}
}

func (l *Logger) Error(msg string) {
	log := &Log{
		Message:   msg,
		Level:     LogError,
		Timestamp: time.Now().UnixMilli(),
		Os:        runtime.GOOS,
	}

	b, err := json.Marshal(log)
	if err != nil {
		// TODO: handle error properly
	}

	if err := l.write(b); err != nil {
		// TODO: handle error properly
	}
}

func (l *Logger) Fatal(msg string) {
	log := &Log{
		Message:   msg,
		Level:     LogFatal,
		Timestamp: time.Now().UnixMilli(),
		Os:        runtime.GOOS,
	}

	b, err := json.Marshal(log)
	if err != nil {
		// TODO: handle error properly
	}

	if err := l.write(b); err != nil {
		// TODO: handle error properly
	}

}

func (l *Logger) write(b []byte) error {
	var err error

	_, err = l.f.Write(b)
	if err != nil {
		return err
	}

	return err
}

func newLogger(ctx context.Context) *Logger {
	return &Logger{ctx: ctx}
}
