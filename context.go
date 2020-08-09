package goworker

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"encoding/xml"
)

type Context struct {
	response http.ResponseWriter
	request *http.Request
	params httprouter.Params
}

func (c *Context) ByName(name string) string{
	return c.params.ByName(name)
}

func (c *Context) BindJson(body []byte , object interface{}) error {
	return json.Unmarshal(body, object)
}

func (c *Context) WriteJson(object interface{}) error {
	body, err := json.Marshal(object)
	if err != nil {
		return err
	}
	c.response.Write(body)
	return err
}

func (c *Context) BindXML(body []byte , object interface{}) error {
	return xml.Unmarshal(body, object)
}

func (c *Context) WriteXML(object interface{}) error {
	body, err := xml.Marshal(object)
	if err != nil {
		return err
	}
	c.response.Write(body)
	return err
}

func (c *Context)WriteStatusCode(code int) error {
	c.response.WriteHeader(code)
	return nil
}

func (c *Context)WriteString(body string) error {
	c.response.Write([]byte(body))
	return nil
}