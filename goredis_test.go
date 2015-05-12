package goredis

import (
	"github.com/alicebob/miniredis"
	"testing"
)

func Test(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	s.RequireAuth("123456")

	addr := s.Addr()

	c := NewClient(addr, "123456")
	defer c.Close()

	conn, err := c.Get()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	if pong, err := String(conn.Do("PING")); err != nil {
		t.Fatal(err)
	} else if pong != "PONG" {
		t.Fatal(pong)
	}

	if pong, err := String(conn.Do("PING")); err != nil {
		t.Fatal(err)
	} else if pong != "PONG" {
		t.Fatal(pong)
	}
}

func TestSelect(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	addr := s.Addr()

	c := NewClient(addr, "")
	defer c.Close()
	c.SetDBIndex(2)

	c2 := NewClient(addr, "")
	defer c2.Close()

	key := "test_key"
	value := "test_value"

	if res, err := String(c.Do("SET", key, value)); err != nil {
		t.Fatal(err)
	} else if res != "OK" {
		t.Fatal("expected to be OK")
	}

	// get from db 2
	if got, err := String(c.Do("GET", key)); err != nil {
		t.Fatal(err)
	} else if got != value {
		t.Fatal(got)
	}

	// get from db 0
	if _, err := String(c2.Do("GET", key)); err != ErrNil {
		t.Fatal(err)
	}
}
