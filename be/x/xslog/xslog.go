package xslog

import (
	"log/slog"
	"os"
)

func FatalErr(err error) {
	slog.Error(err.Error())
	os.Exit(1)
}

func FatalMsg(msg string) {
	slog.Error(msg)
	os.Exit(1)
}
