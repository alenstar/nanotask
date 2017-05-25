package utils

import (
	"crypto/md5"
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	// "encoding/base64"
	"encoding/hex"
)

func reflectCopy(dst, src reflect.Value) error {
	if dst.CanSet() && dst.Kind() == src.Kind() {
		dst.Set(src)
	} else {
		return errors.New("not support")
	}
	return nil
}

func findField(face interface{}, name string) (reflect.Value, error) {
	val := reflect.ValueOf(face).Elem()
	if val.Kind() != reflect.Struct {
		return val, errors.New("not support")
	}
	return val.FieldByName(name), nil
}

func TernaryIf(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func FileLoad(filename string) []byte {
	var file, err = os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		// TODO
		return nil
	}
	defer file.Close()

	n, err := ioutil.ReadAll(file) //file.Read(text)
	if err != nil {
		// TODO
		return nil
	}
	return n
}

func Md5String(str string) string {
	md5ctx := md5.New()
	md5ctx.Write([]byte(str))
	md5encode := md5ctx.Sum(nil)
	return hex.EncodeToString(md5encode)
}
