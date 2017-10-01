package rpc

import (
	"encoding/json"
	"github.com/alenstar/nanotask/log"
	"reflect"
	"fmt"
)

type RPCEcho interface {
	Echo([]byte) ([]byte, error)
}

type Client struct {
	echo RPCEcho
}

// for client to call method
// result must struct pointer
// (args...) --> result
func (c *Client) Call(name string, args interface{}, result interface{}) error {
	req := rpcRequest{
		Version: jsonrpcVersion,
		Method:  name,
		Id:      1, // must greater than zero, on id is zero that this call is notify
		Params:  args,
	}

	body, err := json.Marshal(req)
	if err != nil {
		log.Error("json.Marshal:", err.Error())
		return nil
	}

	reply, err := c.echo.Echo(body)
	if err != nil {
		log.Error("echo.Echo:", err.Error())
		return err
	}
	resp := &rpcResponse{
		Result:result,
	}
	err = json.Unmarshal(reply, resp)
	if err != nil {
		log.Error("json.Unmarshal:", err.Error(), string(reply))
		return err
	} else if resp.Error != nil {
		return fmt.Errorf("RPCError:%s", resp.Error.Error())
	}
	return err
}


// TODO
func (c *Client) Notify(name string, args[]interface{}) error {
	return nil
}

func NewClient(echo RPCEcho) *Client {
	return &Client{echo: echo}
}

func ReflectCopy(dst, src reflect.Value) error {
	if dst.Kind() != src.Kind() {
		if dst.Kind() == reflect.Ptr || dst.Kind() == reflect.Interface {
			return ReflectCopy(dst.Elem(), src)
		} else if src.Kind() == reflect.Ptr || src.Kind() == reflect.Interface {
			return ReflectCopy(dst, src.Elem())
		}
		return fmt.Errorf("not support: %s %s", dst.Kind().String(), src.Kind().String())
	}
	switch dst.Kind() {
	case reflect.Ptr, reflect.Interface:
		return ReflectCopy(dst.Elem(), src.Elem())
	case reflect.Slice:
		if !dst.CanSet() {
			return fmt.Errorf("not set slice")
		}
		dst.Set(reflect.MakeSlice(src.Type(), src.Len(), src.Cap()))
		for i := 0; i < src.Len(); i++ {
			err := ReflectCopy(dst.Index(i), src.Index(i))
			if err != nil {
				return err
			}
		}
	case reflect.Map:
		dst.Set(reflect.MakeMap(src.Type()))
		keys := src.MapKeys()
		if keys != nil {
			for i := 0; i < len(keys); i++ {
				value := src.MapIndex(keys[i])
				dst.SetMapIndex(keys[i], value)
			}
		}
	case reflect.Struct:
		for i := 0; i < dst.NumField(); i++ {
		err := ReflectCopy(dst.Field(i), src.Field(i))
		if err != nil {
			return err
		}
		}
	default: // reflect.Array, reflect.Struct, reflect.Interface
		// not thing todo
		if dst.CanSet() {
			dst.Set(src)
		}
	}

	//	dst.Set(src)
	return nil
}
