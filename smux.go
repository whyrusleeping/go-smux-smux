package peerstream_smux

import (
	"errors"
	"net"

	smux "github.com/jbenet/go-stream-muxer"
	muxer "github.com/xtaci/smux"
)

var ErrUseServe = errors.New("not implemented, use Serve")

type conn struct {
	*muxer.Session
}

func (c *conn) Close() error {
	return c.Session.Close()
}

func (c *conn) IsClosed() bool {
	return c.Session.IsClosed()
}

// OpenStream creates a new stream.
func (c *conn) OpenStream() (smux.Stream, error) {
	return c.Session.OpenStream()
}

// AcceptStream accepts a stream opened by the other side.
func (c *conn) AcceptStream() (smux.Stream, error) {
	return c.Session.AcceptStream()
}

// Serve starts listening for incoming requests and handles them
// using given StreamHandler
func (c *conn) Serve(handler smux.StreamHandler) {
	for {
		s, err := c.AcceptStream()
		if err != nil {
			return
		}
		go handler(s)
	}
}

// Transport is a go-peerstream transport that constructs
// multiplex-backed connections.
type Transport muxer.Config

// DefaultTransport has default settings for multiplex
var DefaultTransport = (*Transport)(muxer.DefaultConfig())

func (t *Transport) NewConn(nc net.Conn, isServer bool) (smux.Conn, error) {
	var c *muxer.Session
	var err error
	if isServer {
		c, err = muxer.Server(nc, (*muxer.Config)(t))
	} else {
		c, err = muxer.Client(nc, (*muxer.Config)(t))
	}
	if err != nil {
		return nil, err
	}

	return &conn{c}, nil
}
