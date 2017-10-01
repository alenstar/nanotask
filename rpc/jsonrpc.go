package rpc

import (
	"fmt"
	"reflect"
)

const jsonrpcVersion  = "2.0"

type rpcRequest struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      uint64      `json:"id"`
	Version string      `json:"jsonrpc"` // 2.0
}

type rpcError struct {
	Code    int         `json:"code"`
	Message string `json:"message"`
	Data    interface{} `json:"data, omitempty"`
}

func (e *rpcError) Error() string {
	return fmt.Sprintf("Code(%d): %v", e.Code, e.Message)
}

type rpcResponse struct {
	Result  interface{} `json:"result, omitempty"`
	Error   *rpcError `json:"error", omitempty`
	Id      uint64      `json:"id"`
	Version string      `json:"jsonrpc"` // 2.0
}

type RPCResult = rpcResponse

type rpcMethod struct {
	method reflect.Method
	value  reflect.Value
}

var (
	methodTable map[string]*rpcMethod
)

func init() {
	methodTable = make(map[string]*rpcMethod)
}

func registerMethods(obj interface{}, name string) error {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() == reflect.Struct {
		if name == "" {
			name = val.Type().Name()
		}
		tp := reflect.TypeOf(obj)
		if tp.NumMethod() == 0 {
			return fmt.Errorf("not method be found")
		}
		for i:=0; i < tp.NumMethod(); i++ {
			if _, ok := methodTable[name+"."+tp.Method(i).Name]; ok {
				return fmt.Errorf("the name %s is registered", name)
			}
			methodTable[name+"."+tp.Method(i).Name] = &rpcMethod{
				method: tp.Method(i),
				value:  reflect.ValueOf(obj),
				// typ: tp,
			}
		}
		return nil
	} else if val.Kind() == reflect.Func {
		if _, ok := methodTable[name]; ok {
			return fmt.Errorf("the name %s is registered", name)
		}
		methodTable[name] = &rpcMethod{
			value: val,
			// typ: reflect.TypeOf(obj),
		}
		return nil
	}

	return fmt.Errorf("invalid type: %s, must struct or function", val.Kind().String())
}

// FIXME
// panic on method params is int (json is float64)
func callMethod(name string, args ...interface{}) (results []interface{}, err error) {
	if fn, ok := methodTable[name]; ok {
		in := make([]reflect.Value, 0)
		var out []reflect.Value
		if fn.method.Func.Kind() == reflect.Invalid {
			if !fn.value.Type().IsVariadic() {
				num := fn.value.Type().NumIn()
				if num != len(args) {
					return nil, fmt.Errorf("The number of parameters does not match")
				}
			}
			for _ ,v:= range args{
				in = append(in, reflect.ValueOf(v))
			}
			out = fn.value.Call(in)
		} else {
			if !fn.method.Func.Type().IsVariadic() {
				num := fn.method.Func.Type().NumIn() - 1
				if num != len(args) {
					return nil, fmt.Errorf("The number of parameters does not match")
				}
			}
			in = append(in, fn.value)
			for _ ,v:= range args {
				in = append(in, reflect.ValueOf(v))
			}
			out = fn.method.Func.Call(in)
		}
		for _, v:= range out {
			results = append(results, v.Interface())
		}
		return
	}
	return nil, fmt.Errorf("the %s method not found", name)
}



