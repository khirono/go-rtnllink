package rtnllink

import (
	"syscall"
)

const (
	IFLA_INFO_UNSPEC = iota
	IFLA_INFO_KIND
	IFLA_INFO_DATA
)

type IfInfomsg syscall.IfInfomsg

func (m IfInfomsg) Len() int {
	return syscall.SizeofIfInfomsg
}

func (m IfInfomsg) Encode(b []byte) (int, error) {
	b[0] = m.Family
	// b[1] = m.X__ifi_pad
	native.PutUint16(b[2:4], m.Type)
	native.PutUint32(b[4:8], uint32(m.Index))
	native.PutUint32(b[8:12], m.Flags)
	native.PutUint32(b[12:16], m.Change)
	return m.Len(), nil
}
