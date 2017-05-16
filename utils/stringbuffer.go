package utils

import (
	"bytes"
)

func integer2Bytes(v interface{}) []byte {
	switch v.(type) {
	case int:
		return reverseInteger2Bytes(uint32(v.(int)))
	case uint:
		return reverseInteger2Bytes(uint32(v.(uint)))
	case int8:
		tmp := make([]byte, 1)
		tmp[0] = byte(v.(int8))
		return tmp
	case uint8:
		tmp := make([]byte, 1)
		tmp[0] = byte(v.(uint8))
		return tmp
	case int16:
		return reverseInteger2Bytes(uint16(v.(int16)))
	case uint16:
		vi := v.(uint16)
		tmp := make([]byte, 2)
		tmp[1] = byte(vi >> 8)
		tmp[0] = byte(vi)
		return tmp
	case int32:
		return reverseInteger2Bytes(uint32(v.(int32)))
	case uint32:
		vi := v.(uint32)
		tmp := make([]byte, 4)
		tmp[3] = byte(vi >> 24)
		tmp[2] = byte(vi >> 16)
		tmp[1] = byte(vi >> 8)
		tmp[0] = byte(vi)
		return tmp
	case int64:
		return reverseInteger2Bytes(uint64(v.(int64)))
	case uint64:
		vi := v.(uint64)
		tmp := make([]byte, 8)
		tmp[7] = byte(vi >> 56)
		tmp[6] = byte(vi >> 48)
		tmp[5] = byte(vi >> 40)
		tmp[4] = byte(vi >> 32)
		tmp[3] = byte(vi >> 24)
		tmp[2] = byte(vi >> 16)
		tmp[1] = byte(vi >> 8)
		tmp[0] = byte(vi)
		return tmp
	default:
		panic("unknown type")
	}
	return nil
}

func reverseInteger2Bytes(v interface{}) []byte {
	switch v.(type) {
	case int:
		return reverseInteger2Bytes(uint32(v.(int)))
	case uint:
		return reverseInteger2Bytes(uint32(v.(uint)))
	case int8:
		tmp := make([]byte, 1)
		tmp[0] = byte(v.(int8))
		return tmp
	case uint8:
		tmp := make([]byte, 1)
		tmp[0] = byte(v.(uint8))
		return tmp
	case int16:
		return reverseInteger2Bytes(uint16(v.(int16)))
	case uint16:
		vi := v.(uint16)
		tmp := make([]byte, 2)
		tmp[0] = byte(vi >> 8)
		tmp[1] = byte(vi)
		return tmp
	case int32:
		return reverseInteger2Bytes(uint32(v.(int32)))
	case uint32:
		vi := v.(uint32)
		tmp := make([]byte, 4)
		tmp[0] = byte(vi >> 24)
		tmp[1] = byte(vi >> 16)
		tmp[2] = byte(vi >> 8)
		tmp[3] = byte(vi)
		return tmp
	case int64:
		return reverseInteger2Bytes(uint64(v.(int64)))
	case uint64:
		vi := v.(uint64)
		tmp := make([]byte, 8)
		tmp[0] = byte(vi >> 56)
		tmp[1] = byte(vi >> 48)
		tmp[2] = byte(vi >> 40)
		tmp[3] = byte(vi >> 32)
		tmp[4] = byte(vi >> 24)
		tmp[5] = byte(vi >> 16)
		tmp[6] = byte(vi >> 8)
		tmp[7] = byte(vi)
		return tmp
	default:
		panic("unknown type")
	}
	return nil
}

// not use encoding/binary
// binary.BigEndian.PutUint16
func reverseInteger(v interface{}) interface{} {
	switch v.(type) {
	case int:
		return int(reverseInteger(uint32(v.(int))).(uint32))
	case uint:
		return uint(reverseInteger(uint32(v.(uint))).(uint32))
	case int8:
		return v
	case uint8:
		return v
	case int16:
		return int16(reverseInteger(uint16(v.(int16))).(uint16))
	case uint16:
		vi := v.(uint16)
		return (vi>>8)&0x00ff | (vi<<8)&0xff00
	case int32:
		return int32(reverseInteger(uint32(v.(int32))).(uint32))
	case uint32:
		vi := v.(uint32)
		return (vi>>24)&0x00ff | (vi>>16)&0x00ff00 | (vi<<16)&0x00ff0000 | (vi<<24)&0xff000000
	case int64:
		return int64(reverseInteger(uint64(v.(int64))).(uint64))
	case uint64:
	default:
		panic("unknown type")
	}
	return v
}

type StringBuffer struct {
	bytes.Buffer
	toNetByteOrder bool
	//endia binary.ByteOrder // binary.BigEndia or binary.LitEndia
}

// type StringBuffer bytes.Buffer

func (s *StringBuffer) Join(v interface{}) *StringBuffer {
	switch v.(type) { // x.(type) only used switch ; if v, ok := x.(TypeA); ok {// TODO}; or v := x.(TypeA) has panic
	// case rune: // rune is int32 alias name
	// 	s.WriteRune(v.(rune))
	case []rune:
		for _, r := range v.([]rune) {
			s.WriteRune(r)
		}
	case string:
		s.WriteString(v.(string))
	case []string:
		for _, str := range v.([]string) {
			s.WriteString(str)
		}
	case byte:
		s.WriteByte(v.(byte))
	case []byte:
		s.Write(v.([]byte))
	case bool:
		if v.(bool) {
			s.WriteByte('1')
		} else {
			s.WriteByte('0')
		}
	case int:
		s.Write(reverseInteger2Bytes(v))
	case uint:
		s.Write(reverseInteger2Bytes(v))
	case int16:
		s.Write(reverseInteger2Bytes(v))
	case uint16:
		s.Write(reverseInteger2Bytes(v))
	case int32: // rune is int32 alias name
		//s.WriteRune(v.(rune))
		s.Write(reverseInteger2Bytes(v))
	case uint32:
		s.Write(reverseInteger2Bytes(v))
	case int64:
		s.Write(reverseInteger2Bytes(v))
	case uint64:
		s.Write(reverseInteger2Bytes(v))
	default:
		panic("unknown type")
	}
	return s
}

func NewStringBuffer(str string) *StringBuffer {
	sb := &StringBuffer{}
	sb.WriteString(str)
	return sb
}
