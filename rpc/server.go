package rpc

import (
	"encoding/json"
	"fmt"
	"github.com/alenstar/nanotask/log"
)

type Server struct {
}

// for server to register method
// callName: [TypeName].[MethodName]
func (s *Server) Register(object interface{}) error {
	return registerMethods(object, "")
}
func (s *Server) RegisterName(name string, object interface{}) error {
	return registerMethods(object, name)
}

func (s *Server) call(req *rpcRequest) (resp *rpcResponse, err error) {
	resp = &rpcResponse{
		Id:      req.Id,
		Version: jsonrpcVersion,
	}
	if _, ok := methodTable[req.Method]; ok {
		var res []interface{}
		if params, ok := req.Params.([]interface{}); ok {
			res, err = callMethod(req.Method, params...)
		} else {
			res, err = callMethod(req.Method, req.Params)
		}
		if err != nil {
			log.Error("callMethod:", err.Error())
			resp.Error = &rpcError{
				// TODO
				// Code
				Message: fmt.Sprintf("[server] callMethod: %s", err.Error()),
			}
			return resp, err
		}
		if res != nil {
			if _, ok := res[len(res)-1].(error); ok {
				resp.Error = &rpcError{
					// TODO
					// Code
					Message: res[len(res)-1].(error).Error(),
				}
				if len(res[:len(res)-2]) == 1 {
					resp.Result = res[0]
				} else {
					resp.Result = res[:len(res)-2]
				}
			} else {
				if len(res) == 1 {
					resp.Result = res[0]
				} else {
					resp.Result = res
				}
			}
		}
	} else {
		resp.Error = &rpcError{
			// TODO
			// Code
			Message: fmt.Sprintf("[server] the method not found: %s", req.Method),
		}
	}
	return resp, nil
}

// request --> response
func (s *Server) Call(body []byte) (out []byte, err error) {
	req := rpcRequest{}
	if err = json.Unmarshal(body, &req); err != nil {
		log.Error("json.Unmarshal:", err.Error())
		return nil, err
	}
	resp, err := s.call(&req)
	if err != nil {
		return nil, err
	}
	out, err = json.Marshal(resp)
	return out, err
}
