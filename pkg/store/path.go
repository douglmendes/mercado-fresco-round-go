package store

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(b), "../..") // "../.." <- Seria o caminho onde esse pkg está até a raiz
	return basepath
}

func PathBuilder(path string) string { // Dai chamando essa função no seu código, é só montar a partir da raíz
	return rootDir() + path
}

func GetPathWithLine() (result string) {
	if _, file, line, ok := runtime.Caller(1); ok {
		result = fmt.Sprintf("%v:%v", file, line)
	}

	return result
}
