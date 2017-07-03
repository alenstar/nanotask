package sockets

import (
	"net"
	"time"
)

type UdpSocket struct {
	conn    *net.UDPConn
	address string
	udpAddr *net.UDPAddr
}

func NewUdpSocket(addr string) *UdpSocket {
	return &UdpSocket{address: addr}
}

func (u *UdpSocket) ReadMessage(body []byte) (int, error) {
	u.conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	nread, _, err := u.conn.ReadFromUDP(body)
	return nread, err
}

func (u *UdpSocket) WriteMessage(body []byte) (int, error) {
	//u.conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
	return u.conn.Write(body)
}

func (u *UdpSocket) Connect(address string) (err error) {
	u.address = address
	u.udpAddr, err = net.ResolveUDPAddr("udp4", address)
	if err == nil {
		u.conn, err = net.DialUDP("udp", nil, u.udpAddr)
	}
	return err
}

func (u *UdpSocket) Listen(address string) (err error) {
	u.address = address
	u.udpAddr, err = net.ResolveUDPAddr("udp4", address)
	if err == nil {
		u.conn, err = net.ListenUDP("udp", u.udpAddr)
	}
	return err
}
