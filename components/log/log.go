package log

import (
    "gopkg.in/op/go-logging.v1"
    "os"
    "path/filepath"
    "time"
)

func NewLogger(logType string) *logging.Logger {
    logName := logType + "-" + time.Now().Format("2006-01-02") + ".log";
    logPath := filepath.Join("log/", logName)

    f, err := logFile(logPath)
    if err != nil {
        panic(err)
    }

    format := logging.MustStringFormatter(`%{level:.4s} %{time:15:04:05.000} %{shortfunc} ▶ %{message}`)

    bk := logging.NewLogBackend(f, "", 0)
    bf := logging.NewBackendFormatter(bk, format)

    bk2 := logging.NewLogBackend(os.Stdout, "", 0)
    bf2 := logging.NewBackendFormatter(bk2, format)
    logging.SetBackend(bf, bf2)

    return logging.MustGetLogger(logType)
}

func logFile(path string) (*os.File, error) {
    return os.OpenFile(path, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
}
