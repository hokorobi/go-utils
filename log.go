package goutils

import (
	"fmt"
	"io"
	"log"
	"os"
)

func PrintTee(m interface{}) {
	f, err := os.OpenFile(GetFilenameSameBase(".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Cannot open log file:" + err.Error())
	}
	defer f.Close()

	log.SetOutput(io.MultiWriter(f, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println(m)
}
func PrintfTee(format string, m ...interface{}) {
	PrintTee(fmt.Sprintf(format, m...))
}

func FatalTee(m ...interface{}) {
	PrintTee(fmt.Sprint(m...))
	os.Exit(1)
}
func FatalfTee(format string, m ...interface{}) {
	PrintTee(fmt.Sprintf(format, m...))
	os.Exit(1)
}
