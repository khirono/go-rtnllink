package rtnllink

import (
	"syscall"

	"github.com/khirono/go-nl"
)

func Create(c *nl.Client, ifname string, attrs ...*nl.Attr) error {
	flags := syscall.NLM_F_CREATE
	flags |= syscall.NLM_F_EXCL
	flags |= syscall.NLM_F_ACK
	req := nl.NewRequest(syscall.RTM_NEWLINK, flags)
	err := req.Append(IfInfomsg{})
	if err != nil {
		return err
	}
	err = req.Append(&nl.Attr{
		Type:  syscall.IFLA_IFNAME,
		Value: nl.AttrString(ifname),
	})
	if err != nil {
		return err
	}
	for _, attr := range attrs {
		err = req.Append(attr)
		if err != nil {
			return err
		}
	}
	_, err = c.Do(req)
	return err
}

func Remove(c *nl.Client, ifname string) error {
	flags := syscall.NLM_F_ACK
	req := nl.NewRequest(syscall.RTM_DELLINK, flags)
	ifindex, err := nl.IfnameToIndex(ifname)
	if err != nil {
		return err
	}
	err = req.Append(IfInfomsg{
		Index: int32(ifindex),
	})
	if err != nil {
		return err
	}
	_, err = c.Do(req)
	return err
}

func Up(c *nl.Client, ifname string) error {
	flags := syscall.NLM_F_ACK
	req := nl.NewRequest(syscall.RTM_NEWLINK, flags)
	ifindex, err := nl.IfnameToIndex(ifname)
	if err != nil {
		return err
	}
	err = req.Append(IfInfomsg{
		Change: syscall.IFF_UP,
		Flags:  syscall.IFF_UP,
		Index:  int32(ifindex),
	})
	if err != nil {
		return err
	}
	_, err = c.Do(req)
	return err
}

func Down(c *nl.Client, ifname string) error {
	flags := syscall.NLM_F_ACK
	req := nl.NewRequest(syscall.RTM_NEWLINK, flags)
	ifindex, err := nl.IfnameToIndex(ifname)
	if err != nil {
		return err
	}
	err = req.Append(IfInfomsg{
		Change: syscall.IFF_UP,
		Index:  int32(ifindex),
	})
	if err != nil {
		return err
	}
	_, err = c.Do(req)
	return err
}
