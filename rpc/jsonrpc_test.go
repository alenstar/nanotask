package rpc

import (
	"bytes"
	"context"
	"fmt"
	"github.com/alenstar/nanotask/log"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func print(s, r string) (string, string) {
	fmt.Println(s + " " + r)
	return s, r
}

type UserInfo struct{
	Name string
	Age int
	Health bool
}

type FnTest struct {
	// status int
}

func (f *FnTest) AddUser(name string) *UserInfo {
	return &UserInfo{name, 34, true}
}
func (f *FnTest) Add(a, b float64) float64{
	return a + b
}
func (f *FnTest) Echo(args ...interface{}) ([]interface{}, error) {
	return args, fmt.Errorf("Echo ok")
}
func (f *FnTest) Say(s, r string) (string, string) {
	fmt.Println(s + " " + r)
	return s, r
}

func TestRegister(t *testing.T) {
	err := registerMethods(print, "print")
	assert.Equal(t, nil, err)
	_, err = callMethod("print", "hello", "world")
	assert.Equal(t, nil, err)
	fn := &FnTest{}
	err = registerMethods(fn, "")
	assert.Equal(t, nil, err)
	out, err := callMethod("FnTest.Say", "hello", "world")
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(out))

	out, err = callMethod("FnTest.Say", "hello", "ok")
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(out))
	assert.Equal(t, "hello", out[0])
	assert.Equal(t, "ok", out[1])


	out, err = callMethod("FnTest.Echo", "hello", "ok", 342, 5.435)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(out))
	assert.Equal(t, "hello", (out[0]).([]interface{})[0])
	assert.Equal(t, "ok", (out[0]).([]interface{})[1])
	assert.Equal(t, 342, (out[0]).([]interface{})[2])
	assert.Equal(t, 5.435, (out[0]).([]interface{})[3])

	err = registerMethods(print, "print")
	assert.NotEqual(t, nil, err)
	err = registerMethods(err, "")
	assert.Equal(t, nil, err)
}

type ClientEcho struct {
}

func (c *ClientEcho) Echo(body []byte) ([]byte, error) {
	resp, err := http.Post("http://127.0.0.1:8888/rpc", "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
func httpserver(t *testing.T) *http.Server {
	rpcSrv := &Server{}
	// err := rpcSrv.Register(&FnTest{})
	// assert.Equal(t, nil, err)

	router := gin.Default()
	router.POST("/rpc", func(c *gin.Context) {
		body, err := c.GetRawData()
		assert.Equal(t, nil, err)
		log.Debug("-> ", string(body))
		reply, err := rpcSrv.Call(body)
		log.Debug("<- ", string(reply))
		assert.Equal(t, nil, err)
		c.Data(200, "application/json", reply)
	})

	srv := &http.Server{
		Addr:    ":8888",
		Handler: router,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Info("listen: %s\n", err)
		}
	}()
	return srv
}

func TestCall(t *testing.T) {
	srv := httpserver(t)
	time.Sleep(time.Second * 1)

	rpcClinet := NewClient(&ClientEcho{})

	user := UserInfo{}
	err := rpcClinet.Call("FnTest.AddUser", "alice", &user)
	assert.Equal(t, nil, err)
	log.Debug("RESULT:", user)

	var sum float64
	err = rpcClinet.Call("FnTest.Add", []int{2, 3}, &sum)
	assert.Equal(t, nil, err)
	assert.Equal(t, 5, int(sum))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
