package goutils

import (
	"os"
	"path/filepath"
)

// https://qiita.com/KemoKemo/items/d135ddc93e6f87008521#comment-7d090bd8afe54df429b9
func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
func GetFilenameSameBase(ext string) string {
	exec, _ := os.Executable()
	return filepath.Join(filepath.Dir(exec), getFileNameWithoutExt(exec)+ext)
}
