package utils


import (
	"errors"
	"reflect"
    "os"
    "io/ioutil"
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
