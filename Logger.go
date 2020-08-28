package main

import (
	tm "github.com/buger/goterm"
)

func InitLogger() {
	tm.Clear()
}
func Log(l ...interface{}) {
	StopListening()
	_, _ = tm.Println(l)
	RestartListening()
	tm.Flush()
	ReadCommand()
}
