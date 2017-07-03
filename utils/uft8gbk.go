package utils

import (
	"bytes"
	"container/list"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"sync"
)

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func GbkToUtf8(s []byte) []byte {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		// log.Error("GbkToUtf8:", e.Error())
		return nil
	}
	return d
}

type PacketQueue struct {
	sync.Mutex
	packet *list.List
}

func (p *PacketQueue) Count() int {
	return p.packet.Len()
}

func (p *PacketQueue) Push(val interface{}) {
	p.Lock()
	p.packet.PushFront(val)
	p.Unlock()
}

func (p *PacketQueue) Pop() interface{} {
	p.Lock()
	e := p.packet.Back()
	if e != nil {
		p.packet.Remove(e)
	}
	p.Unlock()
	if e == nil {
		return nil
	}
	return e.Value
}

func PacketQueueNew() *PacketQueue {
	l := list.New()
	l.Init()
	return &PacketQueue{packet: l}
}

type PackBuffer struct {
	Data   []byte
	Pos    uint
	length uint
}

func (buf *PackBuffer) Len() uint {
	return buf.length
}

func PackBufferNew(size uint) *PackBuffer {
	return &PackBuffer{
		Data:   make([]byte, size),
		Pos:    0,
		length: size,
	}
}
