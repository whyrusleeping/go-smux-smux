package peerstream_smux

import (
	"testing"

	test "github.com/jbenet/go-stream-muxer/test"
)

func TestMultiplexTransport(t *testing.T) {
	test.SubtestAll(t, DefaultTransport)
}
