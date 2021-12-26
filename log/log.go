package log

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
func getFilename(ext string) string {
	exec, _ := os.Executable()
	return filepath.Join(filepath.Dir(exec), getFileNameWithoutExt(exec)+ext)
}
func logt(m interface{}) {
	f, err := os.OpenFile(getFilename(".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Cannot open log file:" + err.Error())
	}
	defer f.Close()

	log.SetOutput(io.MultiWriter(f, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println(m)
}
