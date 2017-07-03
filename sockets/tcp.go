package sockets

import (
	"errors"
	"net"
	"time"
)

type TcpSocket struct {
	conn      *net.TCPConn
	address   string
	tcpAddr   *net.TCPAddr
	connected bool
}

func NewTcpSocket(addr string) *TcpSocket {
	return &TcpSocket{address: addr}
}

func (t *TcpSocket) ReadMessage(body []byte) (int, error) {
	if !t.connected {
		return 0, errors.New("connection not connected")
	}
	t.conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	nread, err := t.conn.Read(body)
	if err != nil {
		if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
			// log.Warn("conn.Read timeout")
			// continue
			return nread, nil
		} else {
			//log.Error("conn.Read:", err.Error())
			t.conn.Close()
			t.connected = false
			return nread, err
		}
	}
	return nread, err
}

func (t *TcpSocket) WriteMessage(body []byte) (int, error) {
	if !t.connected {
		return 0, errors.New("connection not connected")
	}
	//t.conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
	n, err := t.conn.Write(body)
	if err != nil {
		t.conn.Close()
		t.connected = false
	}
	return n, err
}

func (t *TcpSocket) Connect(address string) (err error) {
	if t.connected {
		return errors.New("connection is connected")
	}
	t.address = address
	t.tcpAddr, err = net.ResolveTCPAddr("tcp", address)
	if err == nil {
		t.conn, err = net.DialTCP("tcp", nil, t.tcpAddr)
		if t.conn != nil && err == nil {
			t.conn.SetNoDelay(true) // by default
			t.connected = true
		}
	}
	return err
}

func (t *TcpSocket) Listen(address string) (err error) {
	t.address = address
	t.tcpAddr, err = net.ResolveTCPAddr("tcp", address)
	// if err == nil {
	// 	t.conn, err = net.ListenTCP("tcp", t.tcpAddr)
	// 	if t.conn != nil && err == nil {
	// 		//t.conn.SetNoDelay(true) // by default
	// 		t.connected = true
	// 	}
	// }
	return err
}

func (t *TcpSocket) IsConnected() bool {
	return t.connected
}

func (t *TcpSocket) Disconnect() {
	if t.conn != nil && t.connected == true {
		t.connected = false
	}
}

func (t *TcpSocket) Accept() (net.Conn, error) {
	// TODO
	return nil, nil
}
