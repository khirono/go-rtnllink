package rtnllink

import (
	"sync"
	"syscall"
	"testing"

	"github.com/khirono/go-nl"
)

func TestCreate(t *testing.T) {
	var wg sync.WaitGroup
	mux, err := nl.NewMux()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		mux.Close()
		wg.Wait()
	}()
	wg.Add(1)
	go func() {
		mux.Serve()
		wg.Done()
	}()

	conn, err := nl.Open(syscall.NETLINK_ROUTE)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	c := nl.NewClient(conn, mux)

	peer := nl.Attr{
		Type: VETH_INFO_PEER | syscall.NLA_F_NESTED,
		Value: nl.Encoders{
			IfInfomsg{},
			&nl.Attr{
				Type:  syscall.IFLA_IFNAME,
				Value: nl.AttrString("bar"),
			},
		},
	}
	mtu := &nl.Attr{
		Type:  syscall.IFLA_MTU,
		Value: nl.AttrU32(1400),
	}
	linkinfo := &nl.Attr{
		Type: syscall.IFLA_LINKINFO,
		Value: nl.AttrList{
			{
				Type:  IFLA_INFO_KIND,
				Value: nl.AttrString("veth"),
			},
			{
				Type: IFLA_INFO_DATA,
				Value: nl.AttrList{
					peer,
				},
			},
		},
	}
	err = Create(c, "foo", mtu, linkinfo)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUp(t *testing.T) {
	var wg sync.WaitGroup
	mux, err := nl.NewMux()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		mux.Close()
		wg.Wait()
	}()
	wg.Add(1)
	go func() {
		mux.Serve()
		wg.Done()
	}()

	conn, err := nl.Open(syscall.NETLINK_ROUTE)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	c := nl.NewClient(conn, mux)

	err = Up(c, "foo")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDown(t *testing.T) {
	var wg sync.WaitGroup
	mux, err := nl.NewMux()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		mux.Close()
		wg.Wait()
	}()
	wg.Add(1)
	go func() {
		mux.Serve()
		wg.Done()
	}()

	conn, err := nl.Open(syscall.NETLINK_ROUTE)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	c := nl.NewClient(conn, mux)

	err = Down(c, "foo")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemove(t *testing.T) {
	var wg sync.WaitGroup
	mux, err := nl.NewMux()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		mux.Close()
		wg.Wait()
	}()
	wg.Add(1)
	go func() {
		mux.Serve()
		wg.Done()
	}()

	conn, err := nl.Open(syscall.NETLINK_ROUTE)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	c := nl.NewClient(conn, mux)

	err = Remove(c, "foo")
	if err != nil {
		t.Fatal(err)
	}
}
