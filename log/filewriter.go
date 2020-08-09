package log

import (
	"fmt"
	"goworker/log/queue"
	"os"
	"path"
)

type FileWriter struct {
	file *os.File
	q    *queue.Queue
}

func NewFileWriter(filename string) *FileWriter {
	var f *os.File
	var err error
	if checkFileIsExist(filename) {
		f, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
		if err != nil {
			panic(err)
		}
	} else {
		os.MkdirAll(path.Dir(filename), os.ModePerm)
		f, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, os.FileMode(0666))
		if err != nil {
			panic(err)
		}
	}

	out := &FileWriter{
		file: f,
	}
	out.q = queue.New(func(val interface{}) {
		v := val.([]byte)
		if v != nil {
			_, err := out.file.Write(v)
			if err != nil {
				fmt.Println("Write:", err)
			}
		}
	})
	go out.q.Run()
	return out
}

func (f *FileWriter) Close() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	f.q.Done()
	f.file.Sync()
	f.file.Close()
}

func (f *FileWriter) Write(p []byte) (n int, err error) {
	buf := make([]byte, len(p))
	copy(buf, p)
	f.q.Put(buf)
	return len(p), nil
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
